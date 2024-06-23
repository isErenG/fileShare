// tailwind.config.js
module.exports = {
  content: [
    "./templates/*.{html,js}",
  ],
  theme: {
    extend: {
      colors: {
        'back-color': 'rgb(10, 10, 10)',
        'container-color': 'rgb(31, 31, 31)',
        'third-color': 'rgb(218, 0, 55)',
        'fourth-color': 'rgb(131, 3, 31)',
      },
    },
  },
  plugins: [],
}
