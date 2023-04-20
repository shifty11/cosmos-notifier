/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./src/**/*.rs", "./index.html" ],
  theme: {
    screens: {
        sm: "768px",
        md: "1024px",
        lg: "1280px",
    }
  },
  plugins: [],
}

