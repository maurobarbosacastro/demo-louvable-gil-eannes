import { Injectable } from '@angular/core';
import { EmailTableFields, TemplateInterface } from '@app/shared/interfaces/template.interface';
import { JSONTemplate } from 'angular-email-editor/types';

@Injectable({
    providedIn: 'root'
})
export class EmailMockService {

    constructor() {
    }

    designHTML: string = '<!DOCTYPE HTML PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">\n<html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office">\n<head>\n<!--[if gte mso 9]>\n<xml>\n  <o:OfficeDocumentSettings>\n    <o:AllowPNG/>\n    <o:PixelsPerInch>96</o:PixelsPerInch>\n  </o:OfficeDocumentSettings>\n</xml>\n<![endif]-->\n  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">\n  <meta name="viewport" content="width=device-width, initial-scale=1.0">\n  <meta name="x-apple-disable-message-reformatting">\n  <!--[if !mso]><!--><meta http-equiv="X-UA-Compatible" content="IE=edge"><!--<![endif]-->\n  <title></title>\n  \n    <style type="text/css">\n      \n      @media only screen and (min-width: 520px) {\n        .u-row {\n          width: 500px !important;\n        }\n\n        .u-row .u-col {\n          vertical-align: top;\n        }\n\n        \n            .u-row .u-col-100 {\n              width: 500px !important;\n            }\n          \n      }\n\n      @media only screen and (max-width: 520px) {\n        .u-row-container {\n          max-width: 100% !important;\n          padding-left: 0px !important;\n          padding-right: 0px !important;\n        }\n\n        .u-row {\n          width: 100% !important;\n        }\n\n        .u-row .u-col {\n          display: block !important;\n          width: 100% !important;\n          min-width: 320px !important;\n          max-width: 100% !important;\n        }\n\n        .u-row .u-col > div {\n          margin: 0 auto;\n        }\n\n\n        .u-row .u-col img {\n          max-width: 100% !important;\n        }\n\n}\n    \nbody {\n  margin: 0;\n  padding: 0;\n}\n\ntable,\ntr,\ntd {\n  vertical-align: top;\n  border-collapse: collapse;\n}\n\np {\n  margin: 0;\n}\n\n.ie-container table,\n.mso-container table {\n  table-layout: fixed;\n}\n\n* {\n  line-height: inherit;\n}\n\na[x-apple-data-detectors=\'true\'] {\n  color: inherit !important;\n  text-decoration: none !important;\n}\n\n\n\ntable, td { color: #000000; } #u_body a { color: #0000ee; text-decoration: underline; }\n    </style>\n  \n  \n\n</head>\n\n<body class="clean-body u_body" style="margin: 0;padding: 0;-webkit-text-size-adjust: 100%;background-color: #F7F8F9;color: #000000">\n  <!--[if IE]><div class="ie-container"><![endif]-->\n  <!--[if mso]><div class="mso-container"><![endif]-->\n  <table id="u_body" style="border-collapse: collapse;table-layout: fixed;border-spacing: 0;mso-table-lspace: 0pt;mso-table-rspace: 0pt;vertical-align: top;min-width: 320px;Margin: 0 auto;background-color: #F7F8F9;width:100%" cellpadding="0" cellspacing="0">\n  <tbody>\n  <tr style="vertical-align: top">\n    <td style="word-break: break-word;border-collapse: collapse !important;vertical-align: top">\n    <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td align="center" style="background-color: #F7F8F9;"><![endif]-->\n    \n  \n  \n<div class="u-row-container" style="padding: 0px;background-color: transparent">\n  <div class="u-row" style="margin: 0 auto;min-width: 320px;max-width: 500px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;">\n    <div style="border-collapse: collapse;display: table;width: 100%;height: 100%;background-color: transparent;">\n      <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding: 0px;background-color: transparent;" align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:500px;"><tr style="background-color: transparent;"><![endif]-->\n      \n<!--[if (mso)|(IE)]><td align="center" width="500" style="width: 500px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;" valign="top"><![endif]-->\n<div class="u-col u-col-100" style="max-width: 320px;min-width: 500px;display: table-cell;vertical-align: top;">\n  <div style="height: 100%;width: 100% !important;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">\n  <!--[if (!mso)&(!IE)]><!--><div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;"><!--<![endif]-->\n  \n<table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">\n  <tbody>\n    <tr>\n      <td style="overflow-wrap:break-word;word-break:break-word;padding:10px;font-family:arial,helvetica,sans-serif;" align="left">\n        \n  <div>\n    <!doctype html>\n<html lang="en">\n  <head>\n    <!-- Required meta tags -->\n    <meta charset="utf-8">\n    <meta name="viewport" content="width=device-width, initial-scale=1.0">\n    <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@300&display=swap" rel="stylesheet">\n    <style type="text/css">\n    \timg{margin: auto;display: block;max-width: 100%;}\n      p{font-size: 1.2em;margin: 12px 0;font-family: \'Poppins\', sans-serif;font-weight: 300;}\n      .verify{border-radius:4px;padding: 1.3em;max-width:100%;width: 40%;margin: 0 auto;margin-top: 6em;box-shadow: 0px 0px 5px 0px #b1b1b182}\n      .verify a{background-color:#99cb2c;text-align:center;font-size:18px;padding: .5em 1.2em;text-decoration: none;font-family: \'Poppins\', sans-serif;font-weight: 300;border:1px solid #99cb2c;border-radius:5px;color: #fff;margin: 6px 0;display: block;text-transform: capitalize;}.verify a:hover{background-color: #555;color: #fff;border-color: #555}\n    </style>\n  </head>\n  <body>\n    <div class="verify">\n    \t<img src="https://tagpeak.com/resources/upload/images/10854131658359954691670821458.png" alt="logo">\n      <p>Dear %USERNAME%,</p>\n      <p>Welcome to Tagpeak!</p>\n <p>Your password is %PASSWORD%</p>\n     \n      <p>Thank You,</p>\n      <p><a href="https://tagpeak.com/">tagpeak.com</p>\n    </div>\n    \n  </body>\n</html>\n\n  </div>\n\n      </td>\n    </tr>\n  </tbody>\n</table>\n\n  <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->\n  </div>\n</div>\n<!--[if (mso)|(IE)]></td><![endif]-->\n      <!--[if (mso)|(IE)]></tr></table></td></tr></table><![endif]-->\n    </div>\n  </div>\n  </div>\n  \n\n\n    <!--[if (mso)|(IE)]></td></tr></table><![endif]-->\n    </td>\n  </tr>\n  </tbody>\n  </table>\n  <!--[if mso]></div><![endif]-->\n  <!--[if IE]></div><![endif]-->\n</body>\n\n</html>\n';

    designJSON: JSONTemplate = {
        'counters': {
            'u_column': 1,
            'u_row': 1,
            'u_content_html': 1
        },
        'body': {
            'id': 'S1JIBQejBH',
            'rows': [
                {
                    'id': 'HcjvUctLl0',
                    'cells': [
                        1
                    ],
                    'columns': [
                        {
                            'id': '8qyxHrlOi_',
                            'contents': [
                                {
                                    'id': 'l7Y_BNwpK7',
                                    'type': 'html',
                                    'values': {
                                        'html': '<!doctype html>\n<html lang="en">\n  <head>\n    <!-- Required meta tags -->\n    <meta charset="utf-8">\n    <meta name="viewport" content="width=device-width, initial-scale=1.0">\n    <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@300&display=swap" rel="stylesheet">\n    <style type="text/css">\n    \timg{margin: auto;display: block;max-width: 100%;}\n      p{font-size: 1.2em;margin: 12px 0;font-family: \'Poppins\', sans-serif;font-weight: 300;}\n      .verify{border-radius:4px;padding: 1.3em;max-width:100%;width: 40%;margin: 0 auto;margin-top: 6em;box-shadow: 0px 0px 5px 0px #b1b1b182}\n      .verify a{background-color:#99cb2c;text-align:center;font-size:18px;padding: .5em 1.2em;text-decoration: none;font-family: \'Poppins\', sans-serif;font-weight: 300;border:1px solid #99cb2c;border-radius:5px;color: #fff;margin: 6px 0;display: block;text-transform: capitalize;}.verify a:hover{background-color: #555;color: #fff;border-color: #555}\n    </style>\n  </head>\n  <body>\n    <div class="verify">\n    \t<img src="https://tagpeak.com/resources/upload/images/10854131658359954691670821458.png" alt="logo">\n      <p>Dear %USERNAME%,</p>\n      <p>Welcome to Tagpeak!</p>\n <p>Your password is %PASSWORD%</p>\n     \n      <p>Thank You,</p>\n      <p><a href="https://tagpeak.com/">tagpeak.com</p>\n    </div>\n    \n  </body>\n</html>',
                                        'displayCondition': null,
                                        '_styleGuide': null,
                                        'containerPadding': '10px',
                                        'anchor': '',
                                        '_meta': {
                                            'htmlID': 'u_content_html_1',
                                            'htmlClassNames': 'u_content_html'
                                        },
                                        'selectable': true,
                                        'draggable': true,
                                        'duplicatable': true,
                                        'deletable': true,
                                        'hideable': true
                                    }
                                }
                            ],
                            'values': {
                                'backgroundColor': '',
                                'padding': '0px',
                                'border': {},
                                'borderRadius': '0px',
                                '_meta': {
                                    'htmlID': 'u_column_1',
                                    'htmlClassNames': 'u_column'
                                }
                            }
                        }
                    ],
                    'values': {
                        'displayCondition': null,
                        'columns': false,
                        '_styleGuide': null,
                        'backgroundColor': '',
                        'columnsBackgroundColor': '',
                        'backgroundImage': {
                            'url': '',
                            'fullWidth': true,
                            'repeat': 'no-repeat',
                            'size': 'custom',
                            'position': 'center',
                            'customPosition': [
                                '50%',
                                '50%'
                            ]
                        },
                        'padding': '0px',
                        'anchor': '',
                        'hideDesktop': false,
                        '_meta': {
                            'htmlID': 'u_row_1',
                            'htmlClassNames': 'u_row'
                        },
                        'selectable': true,
                        'draggable': true,
                        'duplicatable': true,
                        'deletable': true,
                        'hideable': true
                    }
                }
            ],
            'headers': [],
            'footers': [],
            'values': {
                '_styleGuide': null,
                'popupPosition': 'center',
                'popupWidth': '600px',
                'popupHeight': 'auto',
                'borderRadius': '10px',
                'contentAlign': 'center',
                'contentVerticalAlign': 'center',
                'contentWidth': '500px',
                'fontFamily': {
                    'label': 'Arial',
                    'value': 'arial,helvetica,sans-serif'
                },
                'textColor': '#000000',
                'popupBackgroundColor': '#FFFFFF',
                'popupBackgroundImage': {
                    'url': '',
                    'fullWidth': true,
                    'repeat': 'no-repeat',
                    'size': 'cover',
                    'position': 'center'
                },
                'popupOverlay_backgroundColor': 'rgba(0, 0, 0, 0.1)',
                'popupCloseButton_position': 'top-right',
                'popupCloseButton_backgroundColor': '#DDDDDD',
                'popupCloseButton_iconColor': '#000000',
                'popupCloseButton_borderRadius': '0px',
                'popupCloseButton_margin': '0px',
                'popupCloseButton_action': {
                    'name': 'close_popup',
                    'attrs': {
                        'onClick': 'document.querySelector(\'.u-popup-container\').style.display = \'none\';'
                    }
                },
                'language': {},
                'backgroundColor': '#F7F8F9',
                'preheaderText': '',
                'linkStyle': {
                    'body': true,
                    'linkColor': '#0000ee',
                    'linkHoverColor': '#0000ee',
                    'linkUnderline': true,
                    'linkHoverUnderline': true
                },
                'backgroundImage': {
                    'url': '',
                    'fullWidth': true,
                    'repeat': 'no-repeat',
                    'size': 'custom',
                    'position': 'center'
                },
                '_meta': {
                    'htmlID': 'u_body',
                    'htmlClassNames': 'u_body'
                }
            }
        },
        'schemaVersion': 17
    };

    data: TemplateInterface[] = [
        {
            id: '1',
            code: 'code',
            name: 'Email TemplateInterface 1',
            templateHtml: this.designHTML,
            templateJson: JSON.stringify(this.designJSON),
            createdAt: new Date(),
            updatedAt: new Date(),
            createdBy: 'admin',
            updatedBy: 'admin'
        }
    ];

    tableData: EmailTableFields[] = [
        {
            id: this.data[0].id,
            name: this.data[0].name
        }
    ];
}
