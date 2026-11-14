import { ImageResponse } from "next/og";

export const SOCIAL_IMAGE_SIZE = {
    width: 1200,
    height: 630,
};

export function createSocialImage() {
    return new ImageResponse(
        <div
            style={{
                width: "100%",
                height: "100%",
                position: "relative",
                display: "flex",
                overflow: "hidden",
                color: "#f8f8ff",
                background: "#080a12",
                fontFamily: "Arial, sans-serif",
            }}
        >
            <div
                style={{
                    position: "absolute",
                    inset: 0,
                    display: "flex",
                    opacity: 0.18,
                    backgroundImage:
                        "linear-gradient(rgba(255,255,255,.09) 1px, transparent 1px), linear-gradient(90deg, rgba(255,255,255,.09) 1px, transparent 1px)",
                    backgroundSize: "48px 48px",
                }}
            />
            <div
                style={{
                    position: "absolute",
                    width: 620,
                    height: 620,
                    left: -230,
                    top: -240,
                    display: "flex",
                    borderRadius: 999,
                    background: "rgba(101,72,255,.38)",
                    filter: "blur(90px)",
                }}
            />
            <div
                style={{
                    position: "absolute",
                    width: 540,
                    height: 540,
                    right: -210,
                    bottom: -250,
                    display: "flex",
                    borderRadius: 999,
                    background: "rgba(35,222,195,.22)",
                    filter: "blur(90px)",
                }}
            />

            <div
                style={{
                    margin: 42,
                    padding: "46px 52px",
                    width: 1116,
                    height: 546,
                    position: "relative",
                    display: "flex",
                    flexDirection: "column",
                    border: "1px solid rgba(255,255,255,.14)",
                    borderRadius: 32,
                    background: "rgba(14,16,27,.83)",
                    boxShadow: "0 30px 80px rgba(0,0,0,.45)",
                }}
            >
                <div style={{ display: "flex", alignItems: "center" }}>
                    <div
                        style={{
                            width: 62,
                            height: 62,
                            display: "flex",
                            alignItems: "center",
                            justifyContent: "center",
                            borderRadius: 18,
                            color: "white",
                            fontSize: 39,
                            fontWeight: 400,
                            background: "linear-gradient(145deg, #7c63ff, #4d31e7)",
                            boxShadow: "0 14px 30px rgba(92,61,226,.38)",
                        }}
                    >
                        N
                    </div>
                    <div style={{ display: "flex", marginLeft: 18, fontSize: 32, fontWeight: 700, letterSpacing: -1 }}>
                        nodepad
                    </div>
                    <div
                        style={{
                            marginLeft: "auto",
                            display: "flex",
                            alignItems: "center",
                            padding: "11px 18px",
                            border: "1px solid rgba(132,111,255,.35)",
                            borderRadius: 999,
                            color: "#b9adff",
                            background: "rgba(108,82,245,.11)",
                            fontSize: 17,
                        }}
                    >
                        <span style={{ width: 8, height: 8, marginRight: 10, display: "flex", borderRadius: 99, background: "#66f0d6" }} />
                        Miễn phí · Không cần đăng nhập
                    </div>
                </div>

                <div style={{ marginTop: 48, display: "flex", flexDirection: "column" }}>
                    <div
                        style={{
                            display: "flex",
                            fontSize: 62,
                            lineHeight: 1.08,
                            fontWeight: 750,
                            letterSpacing: -3,
                            background: "linear-gradient(100deg, #ffffff 12%, #ad9dff 60%, #65e8d4)",
                            backgroundClip: "text",
                            color: "transparent",
                        }}
                    >
                        Ý tưởng đến nhanh.
                    </div>
                    <div style={{ marginTop: 4, display: "flex", fontSize: 62, lineHeight: 1.08, fontWeight: 750, letterSpacing: -3 }}>
                        Ghi lại còn nhanh hơn.
                    </div>
                    <div style={{ marginTop: 25, display: "flex", color: "#9b9fb2", fontSize: 22, lineHeight: 1.5 }}>
                        Viết, lưu và chia sẻ ghi chú ngay trên trình duyệt.
                    </div>
                </div>

                <div style={{ marginTop: "auto", display: "flex", alignItems: "center", color: "#777c91", fontSize: 17 }}>
                    <span style={{ display: "flex", color: "#7fe6d2" }}>nodepad.vulebaolong.com</span>
                    <span style={{ marginLeft: "auto", display: "flex" }}>Riêng tư theo liên kết · Đồng bộ đa thiết bị</span>
                </div>
            </div>
        </div>,
        SOCIAL_IMAGE_SIZE,
    );
}
