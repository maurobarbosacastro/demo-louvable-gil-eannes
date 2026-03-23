import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class UpdateWidgets {
  final BuildContext context;

  UpdateWidgets(this.context);

  static const Color grey = Color(0xFF6C6C6C);
  static const Color textColor = Color(0xFF2C2C2C);

  SnackBar snackBar(VoidCallback onPressed, VoidCallback onCancel) {
    return SnackBar(
      backgroundColor: Colors.white,
      behavior: SnackBarBehavior.floating,
      elevation: 0,
      dismissDirection: DismissDirection.none,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(6)),
      margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 4),
      padding: EdgeInsets.zero,
      content: Stack(
        children: [
          ClipRRect(
            borderRadius: BorderRadius.circular(6),
            child: SizedBox(
              height: 100,
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.stretch,
                children: [
                  Container(width: 12, height: 50, color: Colors.blue),
                  const Padding(
                    padding: EdgeInsets.symmetric(horizontal: 8.0),
                    child: Icon(Icons.info, color: Colors.blue, size: 40),
                  ),
                  Flexible(
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      crossAxisAlignment: CrossAxisAlignment.stretch,
                      children: [
                        Text(
                          "Atualização disponivel",
                          style: GoogleFonts.openSans(color: textColor, fontWeight: FontWeight.w800, fontSize: 14),
                        ),
                        Text(
                          "Existe uma nova atualização disponivel na Store.",
                          style: GoogleFonts.openSans(color: textColor),
                        ),
                      ],
                    ),
                  ),

                  Padding(
                    padding: const EdgeInsets.symmetric(vertical: 32.0, horizontal: 8.0),
                    child: TextButton(
                      style: TextButton.styleFrom(backgroundColor: grey),
                      onPressed: onPressed,
                      child: Text(
                        "Atualizar",
                        style: GoogleFonts.openSans(color: Colors.white),
                      ),
                    ),
                  ),
                ],
              ),
            ),
          ),
          Positioned(
            right: 0,
            top: -0,
            child: IconButton(
              padding: EdgeInsets.zero,
              constraints: const BoxConstraints(),
              splashRadius: 5,
              onPressed: onCancel,
              icon: const Icon(Icons.close, size: 14,),
            ),
          )
        ],
      ),
      duration: const Duration(days: 1),
    );
  }

  //(())final snackbar =
  AlertDialog dialog({required VoidCallback onPressed, required VoidCallback onCancel, required String version}) {
    //final width = context.size?.width;

    return AlertDialog(
      // insetPadding: EdgeInsets.symmetric(horizontal: 16),
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
      contentPadding: EdgeInsets.zero,
      content: ClipRRect(
        borderRadius: BorderRadius.circular(10),
        child: SizedBox(
          height: 350,
          width: 350,
          child: Stack(
            alignment: Alignment.center,
            children: [
              Positioned(
                top: -215,
                child: Container(
                  width: 350,
                  height: 350,
                  alignment: Alignment.center,
                  decoration: const ShapeDecoration(
                    shape: RoundedRectangleBorder(borderRadius: BorderRadius.all(Radius.circular(1000))),
                    color: Color(0x34C7C7C7),
                  ),
                ),
              ),
              Column(
                mainAxisSize: MainAxisSize.min,
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Icon(
                    Icons.info,
                    color: Colors.blue,
                    size: 110,
                  ),
                  const SizedBox(
                    height: 16,
                  ),
                  Text(
                    "Atualização necessária",
                    style: GoogleFonts.openSans(color: textColor, fontWeight: FontWeight.w800, fontSize: 20),
                  ),
                  Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
                    child: Text(
                      "É necessário atualizar a aplicação para a nova versão disponivel ( v$version ) para um bom funcionamento da mesma.",
                      textAlign: TextAlign.center,
                      style: GoogleFonts.openSans(color: textColor),
                    ),
                  ),
                  const SizedBox(
                    height: 8,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      TextButton(
                        onPressed: onCancel,
                        style: TextButton.styleFrom(
                            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
                            shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10), side: const BorderSide(color: grey)),
                            foregroundColor: Colors.grey,
                            backgroundColor: Colors.white),
                        child: Text(
                          "Cancelar",
                          style: GoogleFonts.openSans(color: grey, fontSize: 16),
                        ),
                      ),
                      TextButton(
                        style: TextButton.styleFrom(
                            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
                            shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10), side: const BorderSide(color: Colors.blue)),
                            foregroundColor: Colors.white,
                            backgroundColor: Colors.blue),
                        onPressed: onPressed,
                        child: Text(
                          "Atualizar",
                          style: GoogleFonts.openSans(color: Colors.white, fontSize: 16),
                        ),
                      ),
                    ],
                  ),
                  const SizedBox(height: 16)
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
