import { useState } from "react";
import api from "../api/axios";
import type { Link } from "../models/link";
import Input from "./Input";
import LinkCard from "./LinkCard";
import SubmitButton from "./SubmitButton";

export default function () {
    const [url, setUrl] = useState("");
    const [shortenedUrl, setShortenedUrl] = useState("");
    const [isLoading, setIsLoading] = useState(false);

    const handleSubmit = async (e: React.SubmitEvent) => {
        e.preventDefault();

        setIsLoading(true);

        if (!url.trim()) return;

        try {
            const payload: Link = { url };
            const response = await api.post("shorten", payload);

            setShortenedUrl(response.data.short_link);
        } catch (err) {
            console.error("Error ao encurtar link", err);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="flex flex-col w-full gap-4">
            <form
                onSubmit={handleSubmit}
                className="flex flex-row rounded-lg shadow-sm py-8 px-6 justify-around bg-card w-full"
            >
                <Input
                    value={url}
                    onChange={(e) => setUrl(e.target.value)}
                    isLoading={isLoading}
                />

                <SubmitButton isLoading={isLoading} />
            </form>

            {shortenedUrl && <LinkCard link={shortenedUrl} />}
        </div>
    );
}
