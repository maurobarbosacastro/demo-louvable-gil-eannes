import 'package:flutter/material.dart';

/// A flexible widget for styling multiple placeholders in localized strings.
///
/// Allows individual styling of placeholders within a text, supporting:
/// - Multiple placeholders
/// - Different styles per placeholder
/// - Handling of repeated placeholders
///
/// Usage Examples:
///
/// // Single placeholder styling
/// MultiStyledText(
///   fullText: l10n.reachNextLevel("Premium"),
///   styledPlaceholders: {
///     "Premium": TextStyle(fontSize: 18, color: Colors.blue)
///   }
/// )
///
/// // Multiple placeholder styling
/// MultiStyledText(
///   fullText: l10n.welcomeMessage("John", "\$50"),
///   styledPlaceholders: {
///     "John": TextStyle(color: Colors.green),
///     "\$50": TextStyle(fontWeight: FontWeight.bold)
///   }
/// )
///
/// Supports complex localization scenarios with ease.
class MultiStyledText extends StatelessWidget {
  final String fullText;
  final Map<String, TextStyle> styledPlaceholders;
  final TextStyle? defaultStyle;

  const MultiStyledText({
    super.key,
    required this.fullText,
    required this.styledPlaceholders,
    this.defaultStyle,
  });

  @override
  Widget build(BuildContext context) {
    final effectiveDefaultStyle = defaultStyle ?? DefaultTextStyle.of(context).style;

    if (styledPlaceholders.isEmpty) {
      return Text(fullText, style: effectiveDefaultStyle);
    }

    List<TextSpan> spans = [];
    String remainingText = fullText;
    int currentIndex = 0;

    List<PlaceholderMatch> matches = [];
    for (String placeholder in styledPlaceholders.keys) {
      int index = 0;
      while (true) {
        index = remainingText.indexOf(placeholder, index);
        if (index == -1) break;
        matches.add(PlaceholderMatch(
          placeholder: placeholder,
          start: index,
          end: index + placeholder.length,
        ));
        index += placeholder.length;
      }
    }

    matches.sort((a, b) => a.start.compareTo(b.start));

    for (PlaceholderMatch match in matches) {
      if (match.start > currentIndex) {
        spans.add(TextSpan(
          text: fullText.substring(currentIndex, match.start),
        ));
      }

      spans.add(TextSpan(
        text: match.placeholder,
        style: styledPlaceholders[match.placeholder],
      ));

      currentIndex = match.end;
    }

    if (currentIndex < fullText.length) {
      spans.add(TextSpan(
        text: fullText.substring(currentIndex),
      ));
    }

    return Text.rich(
      TextSpan(
        style: effectiveDefaultStyle,
        children: spans.isEmpty ? [TextSpan(text: fullText)] : spans,
      ),
    );
  }
}

class PlaceholderMatch {
  final String placeholder;
  final int start;
  final int end;

  PlaceholderMatch({
    required this.placeholder,
    required this.start,
    required this.end,
  });
}
