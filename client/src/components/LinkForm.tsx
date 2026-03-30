import Input from "./Input";
import SubmitButton from "./SubmitButton";

export default function () {
    return (
        <div className="flex flex-row rounded-lg shadow-sm py-4 px-6 justify-around bg-card w-full">
            <Input />

            <SubmitButton />
        </div>
    );
}
