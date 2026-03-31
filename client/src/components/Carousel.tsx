import { Shield, Sparkle, Zap } from "lucide-react";
import Card from "./Card";

export default function () {
    return (
        <div className="flex flex-row justify-between w-full">
            <Card
                Icon={Shield}
                title="Confiável"
                text="Url's curtas com nosso sistema otimizado"
            />

            <Card
                Icon={Sparkle}
                title="Simples"
                text="Interface simples focadas no que mais importa"
            />

            <Card
                Icon={Zap}
                title="Rápido"
                text="Seus links funcionam perfeitamente sempre"
            />
        </div>
    );
}
