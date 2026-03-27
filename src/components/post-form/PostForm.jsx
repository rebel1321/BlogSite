import React, { useCallback } from "react";
import { useForm } from "react-hook-form";
import { Button, Input, RTE, Select } from "..";
import appwriteService from "../../appwrite/config";
import { useNavigate } from "react-router-dom";
import { useSelector } from "react-redux";

export default function PostForm({ post }) {
    const { register, handleSubmit, watch, setValue, control, getValues } = useForm({
        defaultValues: {
            title: post?.title || "",
            slug: post?.slug || "",
            content: post?.content || "",
            status: post?.status || "active",
        },
    });

    const navigate = useNavigate();
    const userData = useSelector((state) => state.auth.userData);

    const submit = async (data) => {
        try {
            if (post) {
                // Update an existing post
                // Image is optional on update; backend handles deletion of old image
                const updatedPost = await appwriteService.updatePost(post.slug, {
                    title: data.title,
                    content: data.content,
                    status: data.status,
                    image: data.image, // Pass file array; backend handles it
                });

                if (updatedPost) {
                    navigate(`/post/${updatedPost.slug}`);
                } else {
                    console.error("Failed to update post.");
                }
            } else {
                // Create a new post
                // Image is required for new posts
                if (!data.image || !data.image[0]) {
                    console.error("Featured image is required to create a post.");
                    return;
                }

                const newPost = await appwriteService.createPost({
                    title: data.title,
                    slug: data.slug,
                    content: data.content,
                    status: data.status,
                    image: data.image, // Pass file array; backend handles upload
                    userId: userData.$id,
                });

                if (newPost) {
                    navigate(`/post/${newPost.slug}`);
                } else {
                    console.error("Failed to create post.");
                }
            }
        } catch (error) {
            console.error("An error occurred during submission:", error);
        }
    };

    const slugTransform = useCallback((value) => {
        if (value && typeof value === "string") {
            return value
                .trim()
                .toLowerCase()
                .replace(/[^a-zA-Z\d\s]+/g, "-")
                .replace(/\s+/g, "-");
        }
        return "";
    }, []);

    React.useEffect(() => {
        const subscription = watch((value, { name }) => {
            if (name === "title") {
                setValue("slug", slugTransform(value.title), { shouldValidate: true });
            }
        });

        return () => subscription.unsubscribe();
    }, [watch, slugTransform, setValue]);

    return (
        <form onSubmit={handleSubmit(submit)} className="flex flex-wrap">
            <div className="w-2/3 px-2 text-left text-semi-bold">
                <Input
                    label="Title :"
                    placeholder="Title"
                    className="mb-4"
                    {...register("title", { required: true })}
                />
                <Input
                    label="Slug :"
                    placeholder="Slug"
                    className="mb-4"
                    {...register("slug", { required: true })}
                    onInput={(e) => {
                        setValue("slug", slugTransform(e.currentTarget.value), { shouldValidate: true });
                    }}
                />
                <RTE label="Content :" name="content" control={control} defaultValue={getValues("content")} />
            </div>
            <div className="w-1/3 px-2 text-semi-bold text-left">
                <Input
                    label="Featured Image :"
                    type="file"
                    className="mb-4"
                    accept="image/png, image/jpg, image/jpeg, image/gif"
                    {...register("image", { required: !post })}
                />
                {post && post.imageUrl && (
                    <div className="w-full mb-4">
                        <img
                            src={appwriteService.getFilePreview(post.imageUrl)}
                            alt={post.title}
                            className="rounded-lg"
                        />
                    </div>
                )}
                <Select
                    options={["active", "inactive"]}
                    label="Status"
                    className="mb-4"
                    {...register("status", { required: true })}
                />
                <Button type="submit" bgColor={post ? "bg-green-500" : undefined} className="w-full">
                    {post ? "Update" : "Submit"}
                </Button>
            </div>
        </form>
    );
}
