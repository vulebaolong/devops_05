import Image from "next/image";
import Link from "next/link";
import styles from "./not-found.module.css";

function FloatingNote({ secondary = false }: { secondary?: boolean }) {
    return <div className={secondary ? styles.noteTwo : styles.note} aria-hidden="true">
        <div className={styles.line} />
        <div className={styles.line} />
        <div className={styles.line} />
    </div>;
}

export default function NotFound() {
    return <main className={styles.page}>
        <div className={styles.grid} aria-hidden="true" />
        <div className={styles.glow} aria-hidden="true" />
        <FloatingNote />
        <FloatingNote secondary />

        <section className={styles.card}>
            <Link className={styles.brand} href="/" aria-label="Nodepad — Trang chủ">
                <Image className={styles.logo} src="/logo.png" alt="" width={40} height={40} priority />
                <span>nodepad</span>
            </Link>

            <div className={styles.codeWrap}>
                <p className={styles.code}>404</p>
                <span className={styles.spark} aria-hidden="true" />
            </div>

            <div className={styles.eyebrow}>Trang này đã đi lạc</div>
            <h1 className={styles.title}>Không tìm thấy điều bạn đang tìm</h1>
            <p className={styles.description}>
                Liên kết có thể đã hết hạn, bị thay đổi hoặc trang này chưa từng tồn tại.
                Đừng lo, những ý tưởng mới vẫn đang chờ bạn ở Nodepad.
            </p>

            <div className={styles.actions}>
                <Link className={styles.button} href="/">← Về trang chủ</Link>
                <Link className={styles.secondary} href="/#write">Viết note mới ✦</Link>
            </div>
        </section>
    </main>;
}
