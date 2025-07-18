import React from 'react'
import { Link } from 'react-router-dom'
import appwriteService from '../appwrite/config'

function PostCard({ $id, title, featuredImage }) {
    // Get the direct view URL for the image (no transformations)
    const imageUrl = appwriteService.getFilePreview(featuredImage)

    return (
        <Link to={`/post/${$id}`}>
            <div className="w-full bg-gray-100 rounded-xl p-4 hover:shadow-lg transition">
                <div className="w-full mb-4">
                    {imageUrl ? (
                        <img
                            src={imageUrl}
                            alt={title}
                            className="rounded-xl w-full max-h-60 object-cover"
                        />
                    ) : (
                        <div className="w-full h-48 bg-gray-300 flex items-center justify-center rounded-xl text-gray-600">
                            No Image Available
                        </div>
                    )}
                </div>
                <h2 className="text-xl font-bold text-gray-800">{title}</h2>
            </div>
        </Link>
    )
}

export default PostCard
