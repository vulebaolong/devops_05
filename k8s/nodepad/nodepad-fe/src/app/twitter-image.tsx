import { createSocialImage, SOCIAL_IMAGE_SIZE } from "@/lib/social-image";

export const alt = "Nodepad — Ghi chú online nhanh, riêng tư và miễn phí";
export const size = SOCIAL_IMAGE_SIZE;
export const contentType = "image/png";

export default function TwitterImage() {
    return createSocialImage();
}
