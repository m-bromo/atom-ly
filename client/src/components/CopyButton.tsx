import { Copy } from "lucide-react";

interface props {
    onClick: React.MouseEventHandler<HTMLButtonElement>;
}

export default function ({ onClick }: props) {
    return (
        <button
            onClick={onClick}
            className="flex flex-row bg-neutral-800 px-4 py-2 rounded-lg gap-2 transition-all 300 hover:bg-neutral/90"
        >
            <Copy color="white" />

            <p className="text-neutral-50">Copiar Link</p>
        </button>
    );
}
