interface props {
    isLoading: boolean;
}

export default function ({ isLoading }: props) {
    return (
        <button
            type="submit"
            className="w-fit bg-primary text-white py-2 px-4 rounded-lg transition-all 300 disabled:bg-primary/50"
            disabled={isLoading}
        >
            Encurtar
        </button>
    );
}
