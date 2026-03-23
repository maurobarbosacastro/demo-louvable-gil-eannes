<#import "template.ftl" as layout>
<@layout.registrationLayout displayInfo=true displayMessage=!messagesPerField.existsError('username'); section>

<#if section = "form">

<div class="w-full h-full md:pr-30 md:pl-26 sm:px-10 ">
    <!-- Title --->
    <div class=" mb-5 text-4xl font-normal text-licorice leading-none">
        Forgot Password?
    </div>
    <div class="flex items-baseline mb-10">
        <div class="text-waterloo font-normal">
            No worries, we'll send you reset instructions.
        </div>
    </div>

    <form id="kc-reset-password-form" class="${properties.kcFormClass!}" action="${url.loginAction}"
          method="post">
        <div class="${properties.kcFormGroupClass!}">
            <div class="${properties.kcLabelWrapperClass!}">
                <label for="username"
                       class="${properties.kcLabelClass!} block text-project-waterloo font-normal text-md">
                    <#if !realm.loginWithEmailAllowed>${msg("email")}<#elseif !realm.registrationEmailAsUsername>${msg("email")}<#else>${msg("email")}</#if>
                </label>
            </div>
            <div class="${properties.kcInputWrapperClass!}">
                <input type="text" id="username" name="username"
                       class="${properties.kcInputClass!} bg-inherit
                       w-full pb-2 pt-1 border-b border-x-0 border-t-0 border-brand placeholder:text-waterloo focus:outline-none  focus:border-licorice text-licorice"
                       autofocus
                       placeholder="Your email"
                       value="${(auth.attemptedUsername!'')}"
                       aria-invalid="<#if messagesPerField.existsError('username')>true</#if>" />
                <#if messagesPerField.existsError('username')>
                    <span id="input-error-username" class="${properties.kcInputErrorMessageClass!}"
                          aria-live="polite">
                        ${kcSanitize(messagesPerField.get('username'))?no_esc}
                    </span>
                </#if>
            </div>
        </div>
        <div class="${properties.kcFormGroupClass!} ${properties.kcFormSettingClass!}">

            <div id="kc-form-buttons" class="${properties.kcFormButtonsClass!}">
                <button
                    class="${properties.kcButtonClass!} ${properties.kcButtonPrimaryClass!} ${properties.kcButtonBlockClass!} ${properties.kcButtonLargeClass!}
                    !bg-licorice !text-white text-base rounded mt-4 font-normal	"
                    type="submit">
                    Reset
                </button>
            </div>
        </div>
    </form>
    </#if>


    </@layout.registrationLayout>
