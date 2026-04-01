interface props {
    value: string;
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
    isLoading: boolean;
}

export default function ({ value, onChange, isLoading }: props) {
    return (
        <input
            type="text"
            placeholder="Insira seu Link aqui"
            className="w-4/5 bg-background rounded-lg py-2 border-2 border-neutral outline-none px-4 transition-all duration-300 focus:border-primary focus:ring-2 focus:ring-primary"
            value={value}
            onChange={onChange}
            disabled={isLoading}
        />
    );
}
