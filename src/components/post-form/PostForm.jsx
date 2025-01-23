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
            slug: post?.$id || "",
            content: post?.content || "",
            status: post?.status || "active",
        },
    });

    const navigate = useNavigate();
    const userData = useSelector((state) => state.auth.userData);

    const submit = async (data) => {
        try {
            let file = null;

            // Upload file if provided
            if (data.image && data.image[0]) {
                file = await appwriteService.uploadFile(data.image[0]).catch((error) => {
                    console.error("File upload failed:", error);
                    return null;
                });
            }

            if (post) {
                // Update an existing post
                if (file && post.featuredImage) {
                    await appwriteService.deleteFile(post.featuredImage).catch((error) => {
                        console.error("Failed to delete old file:", error);
                    });
                }

                const updatedPost = await appwriteService.updatePost(post.$id, {
                    ...data,
                    featuredImage: file ? file.$id : post.featuredImage,
                });

                if (updatedPost) {
                    console.log("Post updated successfully:", updatedPost);
                    navigate(`/post/${updatedPost.$id}`);
                } else {
                    console.error("Failed to update post.");
                }
            } else {
                // Create a new post
                if (file) {
                    data.featuredImage = file.$id;
                    const newPost = await appwriteService.createPost({
                        ...data,
                        userId: userData.$id,
                    });

                    if (newPost) {
                        console.log("Post created successfully:", newPost);
                        navigate(`/post/${newPost.$id}`);
                    } else {
                        console.error("Failed to create post.");
                    }
                } else {
                    console.error("File upload failed. Cannot create post without a featured image.");
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
                {post && post.featuredImage && (
                    <div className="w-full mb-4">
                        <img
                            src={appwriteService.getFilePreview(post.featuredImage)}
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
