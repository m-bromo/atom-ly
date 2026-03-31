import type { ElementType } from "react";

interface props {
    Icon: ElementType;
    title: string;
    text: string;
}

export default function ({ Icon, title, text }: props) {
    return (
        <div className="size-52 bg-card flex flex-col items-center gap-4 py-6 rounded-xl transition-all duration-300 hover:-translate-y-1 hover:shadow-2xl">
            <Icon size={32} color="#0052ff" />

            <h1 className="text-center font-semibold text-xl">{title}</h1>

            <p className="text-center">{text}</p>
        </div>
    );
}
