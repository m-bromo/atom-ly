import Carousel from "./components/Carousel";
import Footer from "./components/Footer";
import LinkForm from "./components/LinkForm";
import Navbar from "./components/Navbar";

export default function () {
    return (
        <div className="flex flex-col min-h-screen min-w-screen gap-8 items-center bg-background">
            <Navbar />

            <main className="grow flex flex-col gap-16 items-center w-2xl">
                <h1 className="text-5xl text-center">
                    Uma forma segura <br /> de encurtar seus links
                </h1>

                <LinkForm />

                <Carousel />
            </main>

            <Footer />
        </div>
    );
}
