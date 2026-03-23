<#import "template.ftl" as layout>
<@layout.registrationLayout; section>
    <#if section = "header">
        ${msg("emailLinkIdpTitle", idpDisplayName)}
    <#elseif section = "form">
        <div class=" mb-10 text-4xl font-normal text-licorice leading-none">
            Almost there!
        </div>
        <p id="instruction1" class="instruction text-waterloo">
            ${msg("emailLinkIdp1", idpDisplayName, brokerContext.username, realm.displayName)}
        </p>
        <br>
        <p id="instruction2" class="instruction text-waterloo">
            ${msg("emailLinkIdp2")} <a href="${url.loginAction}" class="underline">${msg("doClickHere")}</a> ${msg("emailLinkIdp3")}
        </p>
        <#--<br>
        <p id="instruction3" class="instruction text-waterloo">
            ${msg("emailLinkIdp4")} <a href="${url.loginAction}" class="underline">${msg("doClickHere")}</a> ${msg("emailLinkIdp5")}
        </p>-->
    </#if>
</@layout.registrationLayout>
