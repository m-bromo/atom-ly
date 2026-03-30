import Input from "./Input";
import SubmitButton from "./SubmitButton";

export default function () {
    return (
        <div className="flex flex-row rounded-lg shadow-sm py-4 px-8 gap-2 w-140 bg-card">
            <Input />

            <SubmitButton />
        </div>
    );
}
