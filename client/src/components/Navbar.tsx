import ThemeButton from "./ThemeButton";
import Title from "./Title";

export default function () {
    return (
        <nav className="flex flex-row justify-between w-screen p-6 bg-card">
            <Title />

            <ThemeButton />
        </nav>
    );
}
