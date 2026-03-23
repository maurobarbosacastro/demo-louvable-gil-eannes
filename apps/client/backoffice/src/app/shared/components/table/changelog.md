# Atlanse Table Changelog @1.1.0

All notable changes to the Atlanse Table component will be documented in this file.

### [1.0.0] - 21-02-20223

#### Added

- Initial release of the Atlanse Table component
- Customizable table using Angular Material that uses MatTable as a basis
- Features like sorting, pagination, and ngTemplateOutlet
- Variable CSS for full customization

### [1.1.0] - 15-03-2023

#### Bug Fixing

- Updated the table to allow usage of sorting (it was not displaying the title);
- Fixed paginator default color
- Fix paginator not working when we have filtering

#### Added

- Added event emits on Sort Change
- Added event emits on Page Change
- Added new variable to interface: colWidth
- Hide paginator in case of no results
- Added "disabled" attribute to the actions

### [1.2.0] - 09-10-2023

#### Bug Fixing

- Fix paginator on page change
- Fix total pages on page change


#### Added

- Add noContent div
- Add ids to all items
- Add function to fetch sort

