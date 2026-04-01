import { Check } from "lucide-react";
import CopyButton from "./CopyButton";

interface props {
    link: string;
}

export default function ({ link }: props) {
    const handleCopy = async () => {
        try {
            await navigator.clipboard.writeText(link);
        } catch (err) {
            console.error("Falha ao copiar o texto", err);
        }
    };

    return (
        <div className="flex flex-row w-full bg-primary/20 p-6 rounded-lg justify-between items-center transition-all 300">
            <div className="flex flex-row gap-4 items-center">
                <div className="flex py-2 px-1 bg-primary/50 rounded-lg">
                    <Check />
                </div>

                <div className="flex flex-col">
                    <p>Seu link está pronto</p>

                    <a href={link} target="_blank" rel="noopener noreferrer">
                        {link}
                    </a>
                </div>
            </div>

            <CopyButton onClick={handleCopy} />
        </div>
    );
}
