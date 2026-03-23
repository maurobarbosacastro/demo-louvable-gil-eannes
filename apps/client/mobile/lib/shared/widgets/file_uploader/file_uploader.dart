import 'package:flutter/material.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class FileUploader extends StatefulWidget {
  final void Function(PlatformFile?)? onUpload;
  final void Function()? onRemove;

  const FileUploader({super.key, this.onUpload, this.onRemove});

  @override
  State<FileUploader> createState() => _FileUploaderState();
}

class _FileUploaderState extends State<FileUploader> {
  bool hasError = false;
  PlatformFile? _selectedFile;

  Future<void> pickFile() async {
    try {
      FilePickerResult? result;

      result = await FilePicker.platform.pickFiles(
        type: FileType.custom,
        allowedExtensions: ['pdf'],
      );

      if (result != null) {
        setState(() {
          _selectedFile = result!.files.single;
        });
      }

      widget.onUpload?.call(_selectedFile);
    } catch (e) {
      setState(() => hasError = true);
    }
  }

  @override
  Widget build(BuildContext context) {
    return _selectedFile == null
        ? Button(
            label: translation(context).upload,
            onPressed: pickFile,
            mode: !hasError ? Mode.outlinedDark : Mode.error,
            icon: SvgPicture.asset(
              Assets.upload,
              colorFilter: hasError
                  ? const ColorFilter.mode(Colors.red, BlendMode.srcIn)
                  : null,
            ),
            fontSize: 13,
            width: 84,
          )
        : Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Flexible(
                child: Container(
                  margin: EdgeInsets.only(right: 0),
                  decoration: BoxDecoration(
                    color: AppColors.youngBlue,
                    borderRadius: BorderRadius.circular(4),
                  ),
                  padding: EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      SvgPicture.asset(Assets.checkUnderlined,
                          colorFilter: const ColorFilter.mode(Colors.white, BlendMode.srcIn), height: 15),
                      const SizedBox(width: 8),
                      Expanded(
                        child: Text(
                          _selectedFile!.name,
                          softWrap: true,
                          overflow: TextOverflow.visible,
                          style: Theme.of(context)
                              .textTheme
                              .titleSmall
                              ?.copyWith(color: Colors.white),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              Container(
                height: 30,
                decoration: BoxDecoration(
                    color: AppColors.youngBlue.withAlpha(51),
                    shape: BoxShape.circle),
                child: IconButton(
                  onPressed: () => setState(() {
                    _selectedFile = null;
                    widget.onRemove?.call();
                  }),
                  icon: SvgPicture.asset(Assets.circleClose),
                  color: AppColors.youngBlue,
                  padding: EdgeInsets.zero,
                  iconSize: 16,
                ),
              )
            ],
          );
  }
}
