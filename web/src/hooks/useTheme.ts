import { useStorage } from "@vueuse/core";
enum Theme {
  Light = "light",
  Dark = "dark",
  System = "system",
}
export const useTheme = () => {
  let localtheme = localStorage.getItem("theme");
  if (!localtheme) {
    localtheme = Theme.Light;
  }

  const theme = useStorage("theme", localtheme);
  const toggleTheme = () => {
    if (theme.value === Theme.Light) {
      theme.value = Theme.Dark;
    } else {
      theme.value = Theme.Light;
    }
  };

  const setTheme = (t: Theme) => {
    theme.value = t;
  };

  return {
    theme,
    toggleTheme,
    setTheme,
    Theme,
  };
};
