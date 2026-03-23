<#--<#outputformat "plainText">
    <#assign requiredActionsText>
        <#if requiredActions??>
            <#list requiredActions>
                <#items as reqActionItem>
                    ${msg("requiredAction.${reqActionItem}")}<#sep>, </#sep>
                </#items>
            </#list>
        </#if>
    </#assign>
</#outputformat>

<#import "template.ftl" as layout>
<@layout.emailLayout>
    ${kcSanitize(msg("executeActionsBodyHtml",link, linkExpiration, realmName, requiredActionsText, linkExpirationFormatter(linkExpiration)))?no_esc}
</@layout.emailLayout>-->

<#outputformat "plainText">
    <#assign requiredActionsText>
        <#if requiredActions??>
            <#list requiredActions>
                <#items as reqActionItem>
                    ${msg("requiredAction.${reqActionItem}")}<#sep>, </#sep>
                </#items>
            </#list>
        </#if>
    </#assign>
</#outputformat>

<#import "template.ftl" as layout>
<@layout.emailLayout>
    <#if requiredActions??>
        <#list requiredActions>
            <#items as reqActionItem>
                <#if reqActionItem == "UPDATE_PASSWORD">
                    <div class="clean-body u_body" style="margin: 0;padding: 0;-webkit-text-size-adjust: 100%;background-color: #ffffff;color: #000000">
                        <table style="border-collapse: collapse;table-layout: fixed;border-spacing: 0;mso-table-lspace: 0pt;mso-table-rspace: 0pt;vertical-align: top;min-width: 320px;Margin: 0 auto;background-color: #ffffff;width:100%" cellpadding="0" cellspacing="0">
                            <tbody>
                                <tr style="vertical-align: top">
                                    <td style="word-break: break-word;border-collapse: collapse !important;vertical-align: top">
                                        <div class="u-row-container" style="padding: 0px;background-color: transparent">
                                            <div class="u-row" style="margin: 0 auto;min-width: 320px;max-width: 850px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;">
                                                <div style="border-collapse: collapse;display: table;width: 100%;height: 100%;background-color: transparent;">
                                                    <div class="u-col u-col-100" style="max-width: 320px;min-width: 850px;display: table-cell;vertical-align: top;">
                                                        <div style="height: 100%;width: 100% !important;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
                                                            <!--[if (!mso)&(!IE)]><!--><div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;"><!--<![endif]-->

                                                                <table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
                                                                    <tbody>
                                                                        <tr>
                                                                            <td style="overflow-wrap:break-word;word-break:break-word;padding:10px;font-family:arial,helvetica,sans-serif;" align="left">

                                                                                <div style="font-size: 23px; line-height: 140%; text-align: left; word-wrap: break-word;">
                                                                                    <p style="line-height: 140%;"><strong>Hi there,</strong></p>
                                                                                </div>

                                                                            </td>
                                                                        </tr>
                                                                    </tbody>
                                                                </table>

                                                                <table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
                                                                    <tbody>
                                                                        <tr>
                                                                            <td style="overflow-wrap:break-word;word-break:break-word;padding:10px;font-family:arial,helvetica,sans-serif;" align="left">

                                                                                <div style="font-size: 13px; line-height: 140%; text-align: left; word-wrap: break-word;">
                                                                                    <p style="line-height: 140%;">We received a request to reset your Tagpeak password.</p>
                                                                                    <p style="line-height: 140%;"> </p>
                                                                                    <p style="line-height: 140%;">To start the process, click the link below:</p>
                                                                                    <p style="line-height: 140%;"> </p>
                                                                                    <p style="line-height: 140%;">👉<a href="${link}" style="text-decoration:none">Reset Password</a></p>
                                                                                    <p style="line-height: 140%;"> </p>
                                                                                    <p style="line-height: 140%;">This link will expire within ${linkExpirationFormatter(linkExpiration)} for security reasons.</p>
                                                                                    <p style="line-height: 140%;"> </p>
                                                                                    <p style="line-height: 140%;">If you didn’t request a password reset, you can safely ignore this message your password will remain unchanged.</p>
                                                                                </div>

                                                                            </td>
                                                                        </tr>
                                                                    </tbody>
                                                                </table>

                                                                <table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
                                                                    <tbody>
                                                                        <tr>
                                                                            <td style="overflow-wrap:break-word;word-break:break-word;padding:32px 10px 10px;font-family:arial,helvetica,sans-serif;" align="left">

                                                                                <div style="font-size: 14px; line-height: 140%; text-align: left; word-wrap: break-word;">
                                                                                    <p style="line-height: 140%;">Thanks,</p>
                                                                                    <p style="line-height: 140%;"></p>
                                                                                    <p style="line-height: 140%;">The Tagpeak team</p>
                                                                                </div>

                                                                            </td>
                                                                        </tr>
                                                                    </tbody>
                                                                </table>

                                                                <table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
                                                                    <tbody>
                                                                        <tr>
                                                                            <td style="overflow-wrap:break-word;word-break:break-word;padding:25px 10px;font-family:arial,helvetica,sans-serif;" align="left">

                                                                                <table height="0px" align="center" border="0" cellpadding="0" cellspacing="0" width="100%" style="border-collapse: collapse;table-layout: fixed;border-spacing: 0;mso-table-lspace: 0pt;mso-table-rspace: 0pt;vertical-align: top;border-top: 1px solid #BBBBBB;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%">
                                                                                    <tbody>
                                                                                        <tr style="vertical-align: top">
                                                                                            <td style="word-break: break-word;border-collapse: collapse !important;vertical-align: top;font-size: 0px;line-height: 0px;mso-line-height-rule: exactly;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%">
                                                                                                <span>&#160;</span>
                                                                                            </td>
                                                                                        </tr>
                                                                                    </tbody>
                                                                                </table>

                                                                            </td>
                                                                        </tr>
                                                                    </tbody>
                                                                </table>

                                                                <table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
                                                                    <tbody>
                                                                        <tr>
                                                                            <td style="overflow-wrap:break-word;word-break:break-word;padding:10px;font-family:arial,helvetica,sans-serif;" align="left">

                                                                                <div style="font-size: 14px; line-height: 140%; text-align: center; word-wrap: break-word;">
                                                                                    <p style="line-height: 140%;"><em><span style="color: #ced4d9; line-height: 19.6px;">Copyright © 2025 Tagpeak, All rights reserved.</span></em></p>
                                                                                    <p style="line-height: 140%;"> </p>
                                                                                    <p style="line-height: 140%;">
                                                                                        <br />
                                                                                        <span style="color: #ced4d9; line-height: 19.6px;">
                                                                                            <strong>Want to change how you receive these emails?</strong>
                                                                                        </span>
                                                                                        <br />
                                                                                        <span style="color: #ced4d9; line-height: 19.6px;">
                                                                                            You can update your
                                                                                            <a  style="color: inherit !important;"
                                                                                                href="https://tagpeak.us2.list-manage.com/profile?u=a6da56b488d2d499b9f5049af&id=7f610c59b0&e=3fe8939913&c=95420d1cfa">preferences</a>
                                                                                            or
                                                                                            <a  style="color: #ced4d9 !important; "
                                                                                                href="https://tagpeak.us2.list-manage.com/profile?u=a6da56b488d2d499b9f5049af&id=7f610c59b0&e=3fe8939913&c=95420d1cfa">unsubscribe</a>
                                                                                            from this list.
                                                                                        </span>
                                                                                    </p>
                                                                                </div>

                                                                            </td>
                                                                        </tr>
                                                                    </tbody>
                                                                </table>

                                                                <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>

                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>

                <#else>
                    ${kcSanitize(msg("executeActionsBodyHtml",link, linkExpiration, realmName, requiredActionsText, linkExpirationFormatter(linkExpiration)))?no_esc}
                </#if>
            </#items>
        </#list>
    </#if>


<#--    ${kcSanitize(msg("executeActionsBodyHtml",link, linkExpiration, realmName, requiredActionsText, linkExpirationFormatter(linkExpiration)))?no_esc}-->
</@layout.emailLayout>
