# Atlanse Table @1.2.0

This is an Angular Standalone Component that provides a customizable table using Angular Material that uses MatTable (https://material.angular.io/components/table/overview) as basis.
It allows you to create custom tables, based on the data provided.

It provides features like sorting, pagination and uses ngTemplateOutlet and variable css to be fully customizable.

## Usage

To use the dynamic table component in your Angular application, you need to import it in your module. You also need to import the required Angular Material
modules.

Here is an example of how to use the component:

<details>
<summary> your.component.module </summary>

Configure your module

```
@NgModule({
    declarations: [
    ],
    imports: [
        TableComponent
    ],
})
```

</details>

<details>
<summary> your.component.html </summary>

```
    <atl-ng-table 
        [tableConfiguration]="tableConfiguration"
    >
    </atl-ng-table>

```

</details>

<details>
<summary> your.component.ts </summary>

```
tableConfiguration: ITableConfiguration;
 dataSource: MatTableDataSource<any> = new MatTableDataSource<any>() -> this expects only what's inside content
```

Example:

```

[{
    content: 'Is this a demo?',
    status: 'Draft',
    sponsor: 'Sponsor',
    level: 'Level 1'
}]
    
```

</details>

## Interface

The component uses multiple components with the following properties:

**ITableConfiguration**

<details>
<summary> Interface Specifications </summary>

- `datasource` (required) : MatTableDataSource
- `pageSize` (required): number - pageSize number
- `css` : string[] - list of css to be passed on the table component (general)
- `columns`(required) - Array of **ColumnInterface** (see below) - column specification
- `styles` - **StylesInterface** (see below) - styles to be used on header/body/paginator
- `expanded` - boolean - use if we want the row to have a custom collapsable sub-row (https://material.angular.io/components/table/examples#table-expandable-rows)

</details>

**IColumn**

<details>
<summary> Interface Specifications </summary>

- `id` (required) - string - id of the column (ex. key of the database column to display)
- `name` (required) - string - translated name to display
- `hasSort` (required) - boolean - true/false if you want this column to be sortable
- `hasTooltip` (required) - boolean - true/false if you want a tooltip to be displayed on hover (row value)
- `sortActionDescription` - string - description on the sort action - similar to tooltip
- `valueTransform` - function(row, value) - can be used to obtain a value on a specific row
- `actions` -  **ActionsInterface** - action buttons - edit/delete/add/...
- `css`- function(row, value) - can be used to trigger a css depending on the row value
- `colWidth` - string - can be used to resize column (ex: 20%)

</details>

**IActions**

<details>
<summary> Interface Specifications </summary>

- `type`(required) - string - type of the action (name to display)
- `iconUrl` (required) - string - url of the icon to display
- `css` - if we want a specific css on this action
-

</details>

**IStyles**

<details>
<summary> Interface Specifications </summary>

- `header` - string - style of the header - class based
- `content` - string - style of the header - class based
- `paginator` - string - style of the header - class based

</details>

### Example

<details>
<summary> Interface Usage Example</summary>

    this.tableConfiguration = {
        dataSource: dataSource,
        pageSize: 5,
        styles: {
            header: 'font-nunitoBlack text-black',
            content: 'font-nunitoRegular text-black',
            paginator: 'font-nunitoRegular text-black'
        },
        columns: [
            {
                id: 'logo',
                name: '',
                hasSort: false,
                hasTooltip: false
            },
            {
                id: 'content',
                name: 'Content Column',
                hasSort: true,
                hasTooltip: false,
                sortActionDescription: 'Content Sort Action Description'
            },
            {
                id: 'status',
                name: 'Status',
                hasSort: true,
                hasTooltip: false,
                css: (row: any, column: any) => {
                    return `inline-flex items-center font-bold text-xs px-2.5 py-0.5 rounded-full tracking-wide bg-green-200 text-green-800 dark:bg-green-600 dark:text-green-50`;
                }
            },
            {
                id: 'actions',
                name: '',
                hasSort: false,
                hasTooltip: false,
                actions: [
                    {
                        type: 'edit',
                        iconUrl: 'edit'
                    },
                    {
                        type: 'submit',
                        iconUrl: 'submit'
                    }
                        ]
            }
                ],
            }

</details>

## Variable CSS

This component is a reusable UI element that makes use of variable CSS to provide customizable styles. It allows you to define CSS variables for various style properties, which can be easily changed
at runtime to customize the appearance of the component.

### Usage

To use this component, simply include the variable-css-component tag in your HTML markup and define the desired CSS variables in your stylesheet. Here is an example:

    atl-ng-table {
        --header-background-color: white;
        --row-background-color: red;
    }

In this example, we define two CSS variables --header-background-color and --row-background-color and set their default values. These variables can be easily modified at runtime using JavaScript or by
simply changing the value in your stylesheet.

### CSS Variables

The following CSS variables are available for customization:

- `--header-background-color` - The background color of the header
- `--header-text-color` - The text color of the header

- `--content-background-color` - The background color of the content
- `--content-text-color` - The text color of the content

- `--row-background-color` - The background color of the row
- `--row-text-color` - The text color of the row

- `--div-paginator-display` - The display of the paginator div (default: flex)
- `--div-paginator-justify` - The alignment of the paginator div (default: center)
- `--div-paginator-text-color` - The background color of the paginator
- `--div-paginator-background` - The text color of the paginator

- ` --even-background-color` - Background color of even rows
- `--odd-background-color`- Background color of the odd rows

### More information

https://www.w3schools.com/css/css3_variables.asp

## How to implement

This component is an Angular directive that makes use of **ngTemplateOutlet** to provide customizable content rendering. It allows you to define a template in your component and pass it as an input to
another component or directive, which will then render the content.
(https://angular.io/api/common/NgTemplateOutlet)

### Usage

To use this component, you need to define a template in your component and then pass it as an input to the ngTemplateOutlet directive. Here is an example:

    <atl-ng-table *ngIf="tableConfiguration"
        [tableConfiguration]="tableConfiguration"
        [bodyCustomTemplate]="bodyCustomTemplate">
    </atl-ng-table>

``<ng-template #bodyCustomTemplate let-column><div>{{ column.name }}</div></ng-template>``

### Available Templates

- `bodyCustomTemplate` - Template of the content of the table - <td>
- `headerCustomTemplate` - Template of the header - <tr>
- `bodyCustomExpandedTemplate` - Template of the expanded row

By default the **atl-table** has this templates, so, if you don't need to use custom, don't pass them.

### Context

The following inputs are available for the ngTemplateOutlet directive:
***let-implicit***

Using ***let-implicit*** we can retrieve data from the context.
Implicit contains row, column, index and rowValue.

### Example of usage:

    <ng-template #bodyCustomTemplate let-implicit>
        <div *ngIf="implicit.column.id === 'logo' && implicit.row.sponsor === 'Sponsored'">
            <mat-icon svgIcon="dollar">
            </mat-icon>
        </div>
    </ng-template>

## Examples

### Example of usage

***Use Case Scenario***: We want to put a Checkbox in the difficulty column and mark checked where Status is Published

    <ng-template #bodyCustomTemplate let-implicit>
         <div *ngIf="implicit.column.id === 'difficulty'">
            <mat-checkbox class="example-margin"
                      [checked]="implicit.row.status === 'Published'"
                      color="warn">
            </mat-checkbox>
        </div>
    </ng-template>

***Use Case Scenario 2***: We want to have an icon where the column is the id and the row value of the sponsor is 'Sponsored'

    <ng-template #bodyCustomTemplate let-implicit>
        <div *ngIf="implicit.column.id === 'logo' && implicit.row.sponsor === 'Sponsored'">
            <mat-icon svgIcon="dollar"></mat-icon>
        </div>
    </ng-template>


***Use Case Scenario 3***: We want to define a color for a specific status (Example: Draft = Gray, Published = green) - defined on column

        css: (row: any, column: any) => {
            if (column.id === 'status') {
                return `status-column status-badge rounded-3xl p-1 status-${row.status.toLowerCase()}`;
            }
            return null;
        }

***Use Case Scenario 4***: I want to define odd/even colors for my table

    In your component.css, add the following:

    atl-ng-table {
        --even-background-color: red;
    }

***Use Case Scenario 5***: I want to add an expanded element to my table row

    In your tableConfiguration, add the following:

        expanded: true

    In your component.html, add the following:

    <ng-template #expanded>
        Your Expanded Element Here
    </ng-template>

     <atl-ng-table 
        [tableConfiguration]="tableConfiguration"
        [bodyCustomExpandedTemplate]="expanded"
    >
    </atl-ng-table>

## Outputs

This component that provides custom outputs to allow communication between components. It allows you to emit custom events from the component, which can be subscribed to by other components or
services

### Outputs

- `onRowClick` - This event will be emitted whenever we click on a row
- `onActionClick` - This event will be emited whenever we click on any action button
- `onSortAndPageChange` - This event will be emited on page sorting or page changing

## Browser Support

This component is compatible with all modern browsers that support Angular and the ngTemplateOutlet directive. It may not work as expected in older browsers that do not support these features.

## License

This component is released under the Atlanse Portugal.
