<#import "template.ftl" as layout>
<@layout.registrationLayout displayMessage=false; section>
    <#--<#if section = "header">
        ${kcSanitize(msg("errorTitle"))?no_esc}
    <#elseif section = "form">
        <div id="kc-error-message">
            <p class="instruction">${kcSanitize(message.summary)?no_esc}</p>
            <#if skipLink??>
            <#else>
                <#if client?? && client.baseUrl?has_content>
                    <p><a id="backToApplication" href="${client.baseUrl}">${kcSanitize(msg("backToApplication"))?no_esc}</a></p>
                </#if>
            </#if>
        </div>
    </#if>-->
    <div class="w-full h-full md:pr-30 md:pl-26 sm:px-10 ">
        <!-- Title --->
        <div class=" mb-5 text-4xl font-normal text-licorice leading-none">
            Something went wrong
        </div>
        <div class="flex items-baseline mb-10">
            <div class="text-waterloo font-normal">
                <p><a id="backToApplication" href="${properties.loginRedirectUrl}">${kcSanitize(msg("backToApplication"))?no_esc}</a></p>
            </div>
        </div>
    </div>

</@layout.registrationLayout>
