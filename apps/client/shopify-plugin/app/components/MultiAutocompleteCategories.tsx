import {LegacyStack, Tag, Autocomplete, InlineError} from '@shopify/polaris';
import {useState, useCallback, useMemo} from 'react';
import type {CategoriesInterface} from '../interfaces/common.interface';

interface Props {
	categories: CategoriesInterface[];
	selectedOptions: string[];
	setSelectedOptions: React.Dispatch<React.SetStateAction<string[]>>;
	height: string;
	onChange?: (value?: any) => void;
	emptyError?: boolean;
}

export default function MultiAutocompleteCategories({
	categories,
	selectedOptions,
	setSelectedOptions,
	height,
	onChange,
	emptyError,
}: Props) {
	const categoriesMemo = useMemo(
		() =>
			categories
				.map((c) => ({label: c.name, value: c.name.toLowerCase()}))
				.sort((a, b) => a.label.localeCompare(b.label)),
		[categories],
	);
	const [inputValue, setInputValue] = useState('');
	const [options, setOptions] = useState(categoriesMemo);

	const updateText = useCallback(
		(value: string) => {
			setInputValue(value);

			if (value === '') {
				setOptions(categoriesMemo);
				return;
			}

			const filterRegex = new RegExp(value, 'i');
			const resultOptions = categoriesMemo.filter((option) => option.label.match(filterRegex));

			setOptions(resultOptions);
		},
		[categoriesMemo],
	);

	const removeTag = useCallback(
		(tag: string) => () => {
			const options = [...selectedOptions];
			options.splice(options.indexOf(tag), 1);
			setSelectedOptions(options);
		},
		[selectedOptions],
	);

	const verticalContentMarkup =
		selectedOptions.length > 0 ? (
			<LegacyStack spacing="extraTight" alignment="center">
				{selectedOptions.map((option) => {
					let tagLabel = '';
					tagLabel = option.replace('_', ' ');
					tagLabel = titleCase(tagLabel);
					return (
						<Tag key={`option${option}`} onRemove={removeTag(option)}>
							{tagLabel}
						</Tag>
					);
				})}
			</LegacyStack>
		) : null;

	const textField = (
		<Autocomplete.TextField
			onChange={updateText}
			label="Categories"
			value={inputValue}
			placeholder="Search for a category..."
			verticalContent={verticalContentMarkup}
			autoComplete="off"
			helpText={<span>Please choose the categories in which your store fits the best</span>}
		/>
	);

	return (
		<div style={{maxHeight: height}}>
			<Autocomplete
				id="autocomplete-categories"
				preferredPosition={'below'}
				allowMultiple
				options={options}
				selected={selectedOptions}
				textField={textField}
				onSelect={(values) => {
					setSelectedOptions(values);
					if (onChange) {
						onChange();
					}
				}}
				listTitle="Suggested Categories"
			/>
			{emptyError && selectedOptions.length === 0 ? (
				<InlineError message="Select at least 1 category" fieldID="autocomplete-categories" />
			) : null}
		</div>
	);

	function titleCase(string: string) {
		return string
			.toLowerCase()
			.split(' ')
			.map((word) => word.replace(word[0], word[0].toUpperCase()))
			.join('');
	}
}
