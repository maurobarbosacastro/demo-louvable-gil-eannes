<#import "template.ftl" as layout>
<#import "password-commons.ftl" as passwordCommons>
<@layout.registrationLayout displayMessage=!messagesPerField.existsError('password','password-confirm'); section>
    <#if section = "header">
    <#--        ${msg("updatePasswordTitle")}-->
    <#elseif section = "form">
        <div class="w-full h-full md:pr-30 md:pl-26 sm:px-10 ">
            <!-- Title --->
            <div class=" mb-5 text-4xl font-normal text-licorice leading-none">
                Reset password
            </div>
            <div class="flex items-baseline mb-10">
                <div class="text-waterloo font-normal">
                    You are logged in automatically from your reset password link. Now you can reset your password from below.
                </div>
            </div>

            <form id="kc-passwd-update-form" class="${properties.kcFormClass!}" action="${url.loginAction}" method="post">
                <div class="${properties.kcFormGroupClass!}">
                    <div class="${properties.kcLabelWrapperClass!}">
                        <label for="password-new" class="${properties.kcLabelClass!}">${msg("passwordNew")}</label>
                    </div>
                    <div class="${properties.kcInputWrapperClass!}">
                        <div class="${properties.kcInputGroup!}">
                            <input type="password" id="password-new" name="password-new" class="${properties.kcInputClass!}"
                                   autofocus autocomplete="new-password"
                                   placeholder="Your New Password"
                                   aria-invalid="<#if messagesPerField.existsError('password','password-confirm')>true</#if>"
                            />
                            <#--<button class="${properties.kcFormPasswordVisibilityButtonClass!}" type="button" aria-label="${msg('showPassword')}"
                                    aria-controls="password-new" data-password-toggle
                                    data-icon-show="${properties.kcFormPasswordVisibilityIconShow!}" data-icon-hide="${properties.kcFormPasswordVisibilityIconHide!}"
                                    data-label-show="${msg('showPassword')}" data-label-hide="${msg('hidePassword')}">
                                <i class="${properties.kcFormPasswordVisibilityIconShow!}" aria-hidden="true"></i>
                            </button>-->
                        </div>

                        <#if messagesPerField.existsError('password')>
                            <span id="input-error-password" class="${properties.kcInputErrorMessageClass!}" aria-live="polite">
                                ${kcSanitize(messagesPerField.get('password'))?no_esc}
                            </span>
                        </#if>
                    </div>
                </div>

                <div class="${properties.kcFormGroupClass!} mt-8">
                    <div class="${properties.kcLabelWrapperClass!}">
                        <label for="password-confirm" class="${properties.kcLabelClass!}">${msg("passwordConfirm")}</label>
                    </div>
                    <div class="${properties.kcInputWrapperClass!}">
                        <div class="${properties.kcInputGroup!}">
                            <input type="password" id="password-confirm" name="password-confirm"
                                   class="${properties.kcInputClass!}"
                                   autocomplete="new-password"
                                   placeholder="Confirm Your new Password"
                                   aria-invalid="<#if messagesPerField.existsError('password-confirm')>true</#if>"
                            />
                            <#--<button class="${properties.kcFormPasswordVisibilityButtonClass!}" type="button" aria-label="${msg('showPassword')}"
                                    aria-controls="password-confirm" data-password-toggle
                                    data-icon-show="${properties.kcFormPasswordVisibilityIconShow!}" data-icon-hide="${properties.kcFormPasswordVisibilityIconHide!}"
                                    data-label-show="${msg('showPassword')}" data-label-hide="${msg('hidePassword')}">
                                <i class="${properties.kcFormPasswordVisibilityIconShow!}" aria-hidden="true"></i>
                            </button>-->
                        </div>

                        <#if messagesPerField.existsError('password-confirm')>
                            <span id="input-error-password-confirm" class="${properties.kcInputErrorMessageClass!}" aria-live="polite">
                                ${kcSanitize(messagesPerField.get('password-confirm'))?no_esc}
                            </span>
                        </#if>

                    </div>
                </div>

                <div class="${properties.kcFormGroupClass!}">
                    <@passwordCommons.logoutOtherSessions/>

                    <div id="kc-form-buttons" class="${properties.kcFormButtonsClass!}">
                        <#if isAppInitiatedAction??>
                            <input class="${properties.kcButtonClass!} ${properties.kcButtonPrimaryClass!} ${properties.kcButtonLargeClass!}" type="submit" value="${msg("doSubmit")}" />
                            <button class="${properties.kcButtonClass!} ${properties.kcButtonDefaultClass!} ${properties.kcButtonLargeClass!}" type="submit" name="cancel-aia"
                                    value="true" />${msg("doCancel")}</button>
                        <#else>

                            <div id="kc-form-buttons" class="${properties.kcFormButtonsClass!}">
                                <button
                                    class="${properties.kcButtonClass!} ${properties.kcButtonPrimaryClass!} ${properties.kcButtonBlockClass!} ${properties.kcButtonLargeClass!}"
                                    type="submit">
                                    Reset password
                                </button>
                            </div>
                        </#if>
                    </div>
                </div>
            </form>

        </div>

        <script type="module" src="${url.resourcesPath}/js/passwordVisibility.js"></script>
    </#if>
</@layout.registrationLayout>
