import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class SearchableDropdown extends StatefulWidget {
  final List<DropdownOption> items;
  final String? hint;
  final Function(String?)? onChanged;
  final String? value;
  final String label;
  final Widget? suffix;
  final bool readOnly;
  final TextEditingController controller;

  const SearchableDropdown({
    super.key,
    required this.items,
    this.hint,
    this.onChanged,
    this.value,
    required this.label,
    this.suffix,
    this.readOnly = false,
    required this.controller
  });

  @override
  State<SearchableDropdown> createState() => _SearchableDropdownState();
}

class _SearchableDropdownState extends State<SearchableDropdown> {
  final FocusNode _focusNode = FocusNode();
  final LayerLink _layerLink = LayerLink();
  OverlayEntry? _overlayEntry;
  List<DropdownOption> _filteredItems = [];
  String? _selectedValue;
  bool _isOpen = false;

  @override
  void initState() {
    super.initState();
    _filteredItems = widget.items;
    _selectedValue = widget.value;
    if (_selectedValue != null && widget.items.isNotEmpty) {
      final selectedItem = widget.items.firstWhere(
            (element) => element.value == _selectedValue!,
        orElse: () => DropdownOption(label: '', value: ''),
      );

      if (selectedItem.value.isNotEmpty) {
        widget.controller.text = selectedItem.label;
      } else {
        _selectedValue = null;
      }
    }
  }

  @override
  void dispose() {
    widget.controller.dispose();
    _focusNode.dispose();
    super.dispose();
  }

  void _toggleOverlay() {
    if (_isOpen) {
      _hideOverlay();
    } else {
      _showOverlay();
    }
  }

  void _showOverlay() {
    setState(() {
      _filteredItems = widget.items;
    });

    _overlayEntry = _createOverlayEntry();
    Overlay.of(context).insert(_overlayEntry!);
    setState(() {
      _isOpen = true;
    });
  }

  void _hideOverlay() {
    _overlayEntry?.remove();
    _overlayEntry = null;
    if (mounted) {
      setState(() {
        _isOpen = false;
      });
    }
  }

  void _filterItems(String query) {
    setState(() {
      if (query.isEmpty) {
        _filteredItems = widget.items;
      } else {
        _filteredItems = widget.items
            .where((item) => item.label.toLowerCase().contains(query.toLowerCase()))
            .toList();
      }
    });
    _overlayEntry?.markNeedsBuild();
  }

  void _selectItem(String item) {
    setState(() {
      _selectedValue = item;
      widget.controller.text = widget.items.firstWhere((element) => element.value == item).label;
    });
    widget.onChanged?.call(item);
    _hideOverlay();
    _focusNode.unfocus();
  }

  OverlayEntry _createOverlayEntry() {
    RenderBox renderBox = context.findRenderObject() as RenderBox;
    var size = renderBox.size;

    return OverlayEntry(
      builder: (context) => Stack(
        children: [
          Positioned.fill(
            child: GestureDetector(
              onTap: _hideOverlay,
              behavior: HitTestBehavior.opaque,
              child: Container(
                color: Colors.transparent,
              ),
            ),
          ),
          Positioned(
            width: size.width,
            child: CompositedTransformFollower(
              link: _layerLink,
              showWhenUnlinked: false,
              offset: Offset(0.0, size.height + 5.0),
              child: Material(
                elevation: 0,
                child: Container(
                  constraints: const BoxConstraints(maxHeight: 200),
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(6),
                    border: Border.all(color: AppColors.waterloo.withAlpha(153), width: 1),
                  ),
                  child: _filteredItems.isEmpty
                      ? Container(
                    padding: const EdgeInsets.all(6),
                    child: Text(
                      translation(context).defaultWithoutResults,
                      style: const TextStyle(color: AppColors.wildSand),
                    ),
                  )
                      : ListView.builder(
                    padding: EdgeInsets.zero,
                    shrinkWrap: true,
                    itemCount: _filteredItems.length,
                    itemBuilder: (context, index) {
                      return InkWell(
                        onTap: () => _selectItem(_filteredItems[index].value),
                        child: Container(
                          padding: const EdgeInsets.only(left: 15, right: 20, top: 10),
                          height: 40,
                          child: Text(
                            _filteredItems[index].label,
                            style: const TextStyle(fontSize: 16),
                          ),
                        ),
                      );
                    },
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return CompositedTransformTarget(
      link: _layerLink,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(widget.label, style: Theme.of(context).textTheme.titleSmall?.copyWith(height: 0.5)),
          TextFormField(
            readOnly: widget.readOnly,
            controller: widget.controller,
            focusNode: _focusNode,
            decoration: InputDecoration(
              hintText: widget.hint ?? translation(context).selectItem,
              hintStyle: Theme.of(context).textTheme.bodyMedium,
              suffixIcon: Row(
                mainAxisSize: MainAxisSize.min,
                mainAxisAlignment: MainAxisAlignment.end,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  widget.suffix ?? const SizedBox.shrink(),
                  GestureDetector(
                    onTap: _toggleOverlay,
                    child: Icon(
                      _isOpen ? Icons.arrow_drop_up : Icons.arrow_drop_down,
                      color: Colors.grey.shade700,
                      size: 20,
                    ),
                  ),
                ],
              ),
              contentPadding: const EdgeInsets.symmetric(vertical: 10.0, horizontal: 0),
              enabledBorder: UnderlineInputBorder(
                borderSide: BorderSide(color: _selectedValue == null || (_selectedValue != null && _selectedValue!.isEmpty) && _selectedValue!.isEmpty ? AppColors.wildSand : AppColors.licorice),
              ),
              focusedBorder: UnderlineInputBorder(
                borderSide: BorderSide(color: _selectedValue == null || (_selectedValue != null && _selectedValue!.isEmpty) && _selectedValue!.isEmpty ? AppColors.wildSand : AppColors.licorice),
              ),
              floatingLabelBehavior: FloatingLabelBehavior.never,
            ),
            onTap: () {
              if (!widget.readOnly && !_isOpen) {
                _showOverlay();
              }
            },
            onChanged: (value) {
              if (!_isOpen) {
                _showOverlay();
              }
              _filterItems(value);
            },
          ),
        ],
      )
    );
  }
}
