import type { ReactNode } from "react";

interface props {
    icon: ReactNode;
    title: string;
    text: string;
}

export default function ({ icon, title, text }: props) {
    return (
        <div className="size-50 bg-card flex flex-col items-center gap-4 py-6 rounded-xl">
            {icon}

            <h1 className="text-center font-semibold text-xl">{title}</h1>

            <p className="text-center">{text}</p>
        </div>
    );
}
