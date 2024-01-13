/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/html/pages/**/*.{html,js}", "./ui/html/layouts/*.{html,js}", "./ui/html/partials/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        "gray-secondary": '#353535',
        "main": '#00ACD7',
      },
    },
  },
  plugins: [],
}

