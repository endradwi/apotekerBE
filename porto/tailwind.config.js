/** @type {import('tailwindcss').Config} */
import daisyui from "daisyui";
export default {
  content: ["./index.html", "./src/**/*.jsx"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [
      {
        mytheme: {
          primary: "#140f8c",
          secondary: "#0B132A",
          neutral: "#FFFFFF",
          info: "#4F5665",
          success: "#00A700",
          warning: "#D00000",
        },
      },
    ],
  },
  plugins: [
    daisyui,
  ],
}

