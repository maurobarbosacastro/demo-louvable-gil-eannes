export interface OptionDropdownInterface {
	label: string;
	value: string;
	iconName?: string;
}

export interface CssDropdownInterface {
	styleHeader: string;
	styleInput?: string;
	styleList: string;
}

export interface IconDropdownInterface {
	iconName: string;
	iconStyle?: string;
	iconColor?: string;
}

export interface IconInputDropdownInterface extends IconDropdownInterface {
	position: 'left' | 'right';
}
