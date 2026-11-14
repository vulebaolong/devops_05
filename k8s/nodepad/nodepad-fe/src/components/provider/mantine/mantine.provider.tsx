import { createTheme, MantineProvider as MantineProviderRoot, MantineTheme } from "@mantine/core";
import { emotionTransform, MantineEmotionProvider } from "@mantine/emotion";
import { ReactNode } from "react";
import { RootStyleRegistry } from "./EmotionRootStyleRegistry";

const theme = createTheme({
    fontFamily: "var(--font-comfortaa), sans-serif",
    headings: {
        fontFamily: "var(--font-comfortaa), sans-serif",
    },
    components: {
        Button: { defaultProps: { size: "compact-xs" } },
        TextInput: { defaultProps: { size: "xs" } },
        Textarea: { defaultProps: { size: "xs" } },
        PasswordInput: { defaultProps: { size: "xs" } },
        Select: {
            defaultProps: { size: "xs" },
            styles: (theme: MantineTheme) => ({
                dropdown: {
                    borderRadius: theme.radius.md,
                },
            }),
        },
        NumberInput: { defaultProps: { size: "xs" } },
        PinInput: { defaultProps: { size: "xs" } },
        Anchor: { defaultProps: { size: "xs" } },
        Combobox: { defaultProps: { size: "xs" } },
        InputBase: { defaultProps: { size: "xs", radius: "md" } },
        InputLabel: { defaultProps: { fz: "xs", c: "dimmed" } },
        Progress: { defaultProps: { size: "xs" } },
        Slider: { defaultProps: { size: "xs" } },
        DateInput: { defaultProps: { size: "xs", radius: "md" } },
    },
});

export default function MantineProvider({ children }: { children: ReactNode }) {
    return (
        <RootStyleRegistry>
            <MantineEmotionProvider>
                <MantineProviderRoot theme={theme} stylesTransform={emotionTransform} defaultColorScheme="dark">
                    {children}
                </MantineProviderRoot>
            </MantineEmotionProvider>
        </RootStyleRegistry>
    );
}
