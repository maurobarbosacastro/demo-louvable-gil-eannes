<#macro registrationLayout bodyClass="" displayInfo=false displayMessage=true displayRequiredFields=false>
    <!DOCTYPE html>
    <html class="${properties.kcHtmlClass!}"<#if realm.internationalizationEnabled> lang="${locale.currentLanguageTag}"</#if>>

        <head>
            <meta charset="utf-8">
            <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
            <meta name="robots" content="noindex, nofollow">

            <#if properties.meta?has_content>
                <#list properties.meta?split(' ') as meta>
                    <meta name="${meta?split('==')[0]}" content="${meta?split('==')[1]}" />
                </#list>
            </#if>
            <title>Tagpeak</title>
            <link rel="icon" href="${url.resourcesPath}/img/favicon.ico" />
            <link href="https://cdn.jsdelivr.net/npm/daisyui@4.12.14/dist/full.min.css" rel="stylesheet" type="text/css" />

            <#if properties.stylesCommon?has_content>
                <#list properties.stylesCommon?split(' ') as style>
                    <link href="${url.resourcesCommonPath}/${style}" rel="stylesheet" />
                </#list>
            </#if>
            <#if properties.styles?has_content>
                <#list properties.styles?split(' ') as style>
                    <link href="${url.resourcesPath}/${style}" rel="stylesheet" />
                </#list>
            </#if>
            <#if properties.scripts?has_content>
                <#list properties.scripts?split(' ') as script>
                    <script src="${url.resourcesPath}/${script}" type="text/javascript"></script>
                </#list>
            </#if>
            <#--<script type="importmap">
                {
                    "imports": {
                        "rfc4648": "${url.resourcesCommonPath}/node_modules/rfc4648/lib/rfc4648.js"
                }
            }
            </script>
            <script src="${url.resourcesPath}/js/menu-button-links.js" type="module"></script>-->
            <#if scripts??>
                <#list scripts as script>
                    <script src="${script}" type="text/javascript"></script>
                </#list>
            </#if>
            <script type="module">
                import { checkCookiesAndSetTimer } from '';

                checkCookiesAndSetTimer(
                    "${url.ssoLoginInOtherTabsUrl?no_esc}"
                );
            </script>

            <link rel="preconnect" href="https://fonts.googleapis.com">
            <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
            <link href="https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;1,100;1,200;1,300;1,400;1,500;1,600;1,700&display=swap"
                  rel="stylesheet">
        </head>

        <body>
            <div class="h-full">
                <div class="flex flex-col md:flex-row items-center md:items-start sm:justify-center md:justify-start flex-auto min-w-0 p-2.5 bg-white rounded-[2.5rem]">
                    <div class="w-1/2 relative sm:hidden rounded-[2rem] md:flex flex-auto overflow-hidden bg-youngBlue dark:border-l flex-col h-[95vh]">
                        <!-- Top section with flex -->
                        <div class="flex flex-row justify-between items-center absolute w-full top-6 px-6 z-10">
                            <div>
                                <img src="${url.resourcesPath}/img/logo_white.svg" width="123" height="30" alt="logo">
                            </div>

                            <div>
                                <a
                                    href="${properties.loginRedirectUrl}"
                                    class="bg-white rounded px-2.5 py-2 cursor-pointer items-center flex text-sm gap-1 leading-none text-licorice ">

                                    <svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <mask id="mask0_1679_19205" style="mask-type:alpha" maskUnits="userSpaceOnUse" x="0" y="0" width="15" height="15">
                                            <rect width="15" height="15" fill="#32334C"/>
                                        </mask>
                                        <g mask="url(#mask0_1679_19205)">
                                            <path d="M8.75 11.25L5 7.5L8.75 3.75L9.625 4.625L6.75 7.5L9.625 10.375L8.75 11.25Z" fill="#32334C"/>
                                        </g>
                                    </svg>


                                    Back to website
                                </a>
                            </div>
                        </div>

                        <!-- Main carousel container -->
                        <div class="overflow-hidden w-full h-full relative" id="carouselContainer">
                            <div class="flex transition-transform duration-[1500ms] ease-in-out h-full w-full" id="carouselSlides">
                                <!-- Slides will be inserted here dynamically -->
                            </div>
                        </div>

                        <!-- Indicators -->
                        <div class="absolute bottom-6 w-full px-6">
                            <div class="flex gap-1" id="carouselIndicators">
                                <!-- Indicators will be inserted here dynamically -->
                            </div>
                        </div>


                    </div>
                    <div id="kc-content"
                         class="flex justify-start h-[90dvh] md:w-full lg:w-1/2 py-8 md:px-4 md:p-16 sm:p-8 rounded-none shadow-none items-center sm:w-full">
                        <div id="kc-content-wrapper" class="w-full">

                            <#-- App-initiated actions should not see warning messages about the need to complete the action -->
                            <#-- during login.                                                                               -->
                            <#if displayMessage && message?has_content && (message.type != 'warning' || !isAppInitiatedAction??)>
                                <div class="hidden alert-${message.type} ${properties.kcAlertClass!} pf-m-<#if message.type = 'error'>danger<#else>${message.type}</#if>">
                                    <div class="pf-c-alert__icon">
                                        <#if message.type = 'success'><span
                                            class="${properties.kcFeedbackSuccessIcon!}"></span></#if>
                                        <#if message.type = 'warning'><span
                                            class="${properties.kcFeedbackWarningIcon!}"></span></#if>
                                        <#if message.type = 'error'><span
                                            class="${properties.kcFeedbackErrorIcon!}"></span></#if>
                                        <#if message.type = 'info'><span class="${properties.kcFeedbackInfoIcon!}"></span></#if>
                                    </div>
                                    <span class="${properties.kcAlertTitleClass!}">${kcSanitize(message.summary)?no_esc}</span>
                                </div>
                            </#if>

                            <#nested "form">

                            <#--<#if auth?has_content && auth.showTryAnotherWayLink()>
                                <form id="kc-select-try-another-way-form" action="${url.loginAction}" method="post">
                                    <div class="${properties.kcFormGroupClass!}">
                                        <input type="hidden" name="tryAnotherWay" value="on" />
                                        <a href="#" id="try-another-way"
                                           onclick="document.forms['kc-select-try-another-way-form'].submit();return false;">${msg("doTryAnotherWay")}</a>
                                    </div>
                                </form>
                            </#if>-->

<#--                            <#nested "socialProviders">-->

                            <#if displayInfo>
                                <div id="kc-info" class="${properties.kcSignUpClass!}">
                                    <div id="kc-info-wrapper" class="${properties.kcInfoAreaWrapperClass!}">
                                        <#nested "info">
                                    </div>
                                </div>
                            </#if>
                        </div>
                    </div>

                </div>
            </div>
            <script src="${url.resourcesPath}/scripts/carousel.js" type="text/javascript"></script>
        </body>


    </html>
</#macro>
