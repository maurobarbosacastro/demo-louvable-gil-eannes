import 'dart:io';

import 'package:tagpeak/features/auth/screens/register/widgets/circular_button_widget.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:image_picker/image_picker.dart';

class ImagePickerComponent extends StatelessWidget {
  final XFile? image;
  final Function(ImageSource imageSource) getImage;
  final bool canEdit;

  const ImagePickerComponent({super.key, this.image, required this.getImage, this.canEdit = true});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: SizedBox(
        height: 120,
        width: 120,
        child: Stack(
          alignment: Alignment.center,
          children: [
            image == null
                ? const CircleAvatar(
                    radius: 55,
                    backgroundImage: AssetImage(Assets.userDefaultImage),
                    backgroundColor: Colors.transparent,
                  )
                : CircleAvatar(
                    radius: 55,
                    backgroundImage: FileImage(File(image!.path)),
                    backgroundColor: Colors.transparent,
                  ),
            canEdit
                ? Positioned(
                    bottom: 0,
                    right: 0,
                    child: IconButton(
                      splashRadius: 24,
                      onPressed: () {
                        showModalBottomSheet(
                          constraints: const BoxConstraints(
                            maxHeight: 140,
                          ),
                          context: context,
                          builder: (context) => Column(
                            crossAxisAlignment: CrossAxisAlignment.stretch,
                            mainAxisAlignment: MainAxisAlignment.spaceAround,
                            children: [
                              Row(
                                mainAxisAlignment: MainAxisAlignment.spaceAround,
                                children: [
                                  CircularButton(
                                    onTap: () => getImage(ImageSource.camera),
                                    icon: const Icon(Icons.camera, color: Colors.grey, size: 24),
                                    label: AppLocalizations.of(context)!.camera,
                                  ),
                                  CircularButton(
                                    onTap: () => getImage(ImageSource.gallery),
                                    icon: const Icon(Icons.image_search, color: Colors.grey, size: 24),
                                    label: AppLocalizations.of(context)!.gallery,
                                  )
                                ],
                              ),
                            ],
                          ),
                        );
                      },
                      icon: Container(
                        alignment: Alignment.center,
                        width: 32,
                        height: 32,
                        decoration: BoxDecoration(
                            color: Colors.white,
                            borderRadius: const BorderRadius.all(
                              Radius.circular(45),
                            ),
                            border: Border.all(color: Colors.grey, width: 2)),
                        child: const Icon(
                          Icons.edit,
                          size: 20,
                        ),
                      ),
                    ),
                  )
                : const SizedBox(),
          ],
        ),
      ),
    );
  }
}
