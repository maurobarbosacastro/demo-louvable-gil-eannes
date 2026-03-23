import {LegacyStack, Tag, Autocomplete, InlineError} from '@shopify/polaris';
import {useState, useCallback, useMemo} from 'react';
import type {CountryInterface} from '../interfaces/common.interface';

interface Props {
	countries: CountryInterface[];
	selectedOptions: string[];
	setSelectedOptions: React.Dispatch<React.SetStateAction<string[]>>;
	height: string;
	onChange?: (value?: any) => void;
	emptyError?: boolean;
}

export default function MultiAutocompleteCountries({
	countries,
	selectedOptions,
	setSelectedOptions,
	height,
	onChange,
	emptyError,
}: Props) {
	const countriesMemo = useMemo(
		() =>
			countries
				.map((c) => ({label: c.name, value: c.name.toLowerCase()}))
				.sort((a, b) => a.label.localeCompare(b.label)),
		[countries],
	);
	const [inputValue, setInputValue] = useState('');
	const [options, setOptions] = useState(countriesMemo);

	const updateText = useCallback(
		(value: string) => {
			setInputValue(value);

			if (value === '') {
				setOptions(countriesMemo);
				return;
			}

			const filterRegex = new RegExp(value, 'i');
			const resultOptions = countriesMemo.filter((option) => option.label.match(filterRegex));

			setOptions(resultOptions);
		},
		[countriesMemo],
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
			label="Countries"
			value={inputValue}
			placeholder="Search for a country..."
			verticalContent={verticalContentMarkup}
			autoComplete="off"
			helpText={<span>Please select all countries where you currently ship to</span>}
		/>
	);

	return (
		<div style={{maxHeight: height}}>
			<Autocomplete
				id="autocomplete-countries"
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
				listTitle="Suggested Countries"
			/>
			{emptyError && selectedOptions.length === 0 ? (
				<InlineError message="Select at least 1 country" fieldID="autocomplete-categories" />
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
