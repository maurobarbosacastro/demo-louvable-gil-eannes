<#import "template.ftl" as layout>
<@layout.registrationLayout; section>
    <#if section = "header">
        ${msg("confirmLinkIdpTitle")}
    <#elseif section = "form">
        <form id="kc-register-form" action="${url.loginAction}" method="post" class="invisible">
            <div class="${properties.kcFormGroupClass!}">
<#--                <button type="submit" class="${properties.kcButtonClass!} ${properties.kcButtonDefaultClass!} ${properties.kcButtonBlockClass!} ${properties.kcButtonLargeClass!}" name="submitAction" id="updateProfile" value="updateProfile">${msg("confirmLinkIdpReviewProfile")}</button>-->
                <button type="submit" class="${properties.kcButtonClass!} ${properties.kcButtonDefaultClass!} ${properties.kcButtonBlockClass!} ${properties.kcButtonLargeClass!}" name="submitAction" id="linkAccount" value="linkAccount">${msg("confirmLinkIdpContinue", idpDisplayName)}</button>
            </div>
        </form>
        <div class="flex justify-center">
            <span class="loader"></span>
        </div>
        <script>
            document.getElementById("linkAccount").click();
        </script>
    </#if>
</@layout.registrationLayout>
