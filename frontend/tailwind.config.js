/** @type {import('tailwindcss').Config} */

export default {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  theme: {
    extend: {
         colors: {
            'cl1': '#14110F',
            'cl2': '#34312D',
            'cl3': '#7E7F83',
            'cl4': '#D9C5B2',
            'cl5': '#F3F3F4',
            'cl6': '#3185FC',
            'cl7': '#313338',
            'cl8': '#9da1a8',
            'cl9': '#eeeff2',
            'cl10': '#0D1117',
            'cl11': '#010409',
            'cl12': '#5865f2',
            'cl13': '#67c0c0',
            'cl14': '#f74828'
        }
    },
  },
  plugins: [],
}

