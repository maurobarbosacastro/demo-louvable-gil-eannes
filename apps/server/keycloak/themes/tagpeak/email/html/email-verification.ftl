<#-- Dispatcher in email-verification.ftl -->

<#assign attrs = user.attributes!{}>

<#-- Get attribute as string, whether it's a sequence or a plain string -->
<#function attr name>
    <#assign v = attrs[name]!>
    <#if v?is_sequence>
        <#return (v?first!'')?string>
    <#elseif v?is_string>
        <#return v>
    <#else>
        <#return ''>
    </#if>
</#function>




<#-- Read attributes -->
<#assign source = attr('source')>
<#assign currencySelectedRaw = attr('currency_selected')?lower_case>
<#assign userId = user.id!''>
<#assign emailExtrasRaw = attr('email_extras')>

<#-- Parse email_extras JSON to get store name -->
<#assign storeName = ''>
<#if emailExtrasRaw != ''>
    <#attempt>
        <#assign emailExtras = emailExtrasRaw?eval>
        <#if emailExtras.storeName??>
            <#assign storeName = emailExtras.storeName>
        </#if>
    <#recover>
        <#-- If JSON parsing fails, storeName remains empty -->
    </#attempt>
</#if>

<#-- Interpret currency_selected as boolean true if "true"/"1"/"yes" -->
<#assign currencySelected = ['true','1','yes']?seq_contains(currencySelectedRaw)>

<#-- Branch: Source == "Shopify" AND currency not selected -->
<#if source == 'Shopify' && !currencySelected>
    <#include "email-verification-shopify.ftl">
<#else>
    <#include "email-verification-default.ftl">
</#if>
