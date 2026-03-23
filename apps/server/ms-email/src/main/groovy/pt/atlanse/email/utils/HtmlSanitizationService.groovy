package pt.atlanse.email.utils

import jakarta.inject.Singleton


import org.jsoup.Jsoup
import org.jsoup.safety.Safelist

@Singleton
class HtmlSanitizationService {

    // This method will sanitize HTML input
    String sanitize(String htmlInput) {
        // Define a whitelist (also known as a safeList)
        Safelist safeList = Safelist.relaxed()
            .addTags("html", "head", "body","table", "tbody", "tr", "td", "th", "style", "img", "span", "div") // Add any specific tags missing
            .addAttributes(":all", "style", "align", "border", "cellpadding", "cellspacing", "width", "height", "valign", "src", "alt", "title", "class")
            .addProtocols("img", "src", "http", "https", "data") // Allow images with data URLs too
            .preserveRelativeLinks(true);

        // Clean the input HTML
        return Jsoup.clean(htmlInput, safeList);
    }
}
