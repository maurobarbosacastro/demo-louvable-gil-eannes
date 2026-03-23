<#macro emailLayout>
<html>
    <head>
        <style type="text/css">

            @media only screen and (min-width: 870px) {
                .u-row {
                    width: 850px !important;
                }

                .u-row .u-col {
                    vertical-align: top;
                }


                .u-row .u-col-100 {
                    width: 850px !important;
                }

            }

            @media only screen and (max-width: 870px) {
                .u-row-container {
                    max-width: 100% !important;
                    padding-left: 0px !important;
                    padding-right: 0px !important;
                }

                .u-row {
                    width: 100% !important;
                }

                .u-row .u-col {
                    display: block !important;
                    width: 100% !important;
                    min-width: 320px !important;
                    max-width: 100% !important;
                }

                .u-row .u-col > div {
                    margin: 0 auto;
                }


                .u-row .u-col img {
                    max-width: 100% !important;
                }

            }

            body{margin:0;padding:0}table,td,tr{border-collapse:collapse;vertical-align:top}p{margin:0}.ie-container table,.mso-container table{table-layout:fixed}*{line-height:inherit}a[x-apple-data-detectors=true]{color:inherit!important;text-decoration:none!important}


            table, td { color: #000000; } </style>
    </head>
<body>
    <#nested>
</body>
</html>
</#macro>
