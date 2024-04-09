/** @type {import('tailwindcss').Config} */
module.exports = {
  // content: ["./*.html", "./src/**/*.{css,js,vue}"],
  content: ["./templates/*.html", "./static/*.{css,js}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
  darkMode: 'selector',
}

