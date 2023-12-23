/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.html"],
  theme: {
    extend: {
      "container": {
        "center": true,
        "screens": {
          1200: "1024px"
        },
      
      },
      "fontFamily": {
        "sans": ["Roboto"]
      }
    },
  },
  plugins: [],
}

