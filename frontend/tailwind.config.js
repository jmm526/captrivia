/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/pages/**/*.{js,ts,jsx,tsx,mdx}"],
  theme: {
    extend: {},
    colors: {
      slate: {
        50: "#f9fafb",
        100: "#f4f5f7",
        200: "#e5e7eb",
        300: "#d2d6dc",
        400: "#9fa6b2",
        500: "#6b7280",
        600: "#4b5563",
        700: "#374151",
        800: "#252f3f",
        900: "#161e2e",
      },
    },
  },
  plugins: [],
};
