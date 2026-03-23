tailwind.config = {
    corePlugins: {
        preflight: false
    },
    theme: {
        extend: {
            colors: {
                brand: '#F6F3F4',
                splashScreen: '#F6F3F4',
                sideMenu: '#F6F3F4',
                licorice: '#32334C',
                lavender: '#9C98D1',
                waterloo:"#7A7B92",
                youngBlue: "#4237DA",
                red:"#F2295B",
                grandis:"#FDCA65",
	              wildsand: '#F6F3F4',
            },
            borderRadius: {
                '5': '5px',
                '14': '14px'
            },
            borderWidth:{
                '1.5': '1.5px'
            },
            fontSize: {
                '22': '22px'
            },
            lineClamp: {
                10: '10'
            }
        }
    },
    daisyui: {
        themes: false, // true: all themes | false: only light + dark | array: specific themes like this ["light",
        // "dark", "cupcake"]
        darkTheme: "off", // name of one of the included themes for dark mode
        base: false, // applies background color and foreground color for root element by default
        styled: true, // include daisyUI colors and design decisions for all components
        utils: true, // adds responsive and modifier utility classes
        rtl: false, // rotate style direction from left-to-right to right-to-left. You also need to add dir="rtl" to your html tag and install `tailwindcss-flip` plugin for Tailwind CSS.
        prefix: "", // prefix for daisyUI classnames (components, modifiers and responsive class names. Not colors)
        logs: true, // Shows info about daisyUI version and used config in the console when building your CSS
    }
}
