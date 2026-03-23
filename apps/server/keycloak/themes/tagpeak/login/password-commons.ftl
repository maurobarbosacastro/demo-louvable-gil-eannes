<#macro logoutOtherSessions>
    <div id="kc-form-options" class="${properties.kcFormOptionsClass!} hidden">
        <div class="${properties.kcFormOptionsWrapperClass!}">
            <div class="checkbox w-full">
                <label>
                    <input type="checkbox" id="logout-sessions" name="logout-sessions" value="on" checked>
                    ${msg("logoutOtherSessions")}
                </label>
            </div>
        </div>
    </div>
</#macro>
