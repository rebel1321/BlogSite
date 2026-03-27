import React, {useState, useEffect} from 'react'
import { Container, PostCard } from '../components'
import appwriteService from "../appwrite/config";
import { useSelector } from 'react-redux';

function MyPosts() {
    const [posts, setPosts] = useState([])
    const userData = useSelector((state) => state.auth.userData);

    useEffect(() => {
        if (userData) {
            appwriteService.getPosts([]).then((allPosts) => {
                // Filter posts to show only current user's posts
                const userPosts = Array.isArray(allPosts) 
                    ? allPosts.filter(post => post.userId === userData.email)
                    : [];
                setPosts(userPosts);
            });
        }
    }, [userData])
    
  return (
    <div className='w-full py-8'>
        <Container>
            <h1 className='text-3xl font-bold mb-8'>My Posts</h1>
            {posts.length === 0 ? (
                <div className='text-center text-gray-600'>
                    <p className='text-lg'>You haven't created any posts yet.</p>
                </div>
            ) : (
                <div className='flex flex-wrap'>
                    {posts.map((post) => (
                        <div key={post.slug} className='p-2 w-1/4'>
                            <PostCard {...post} />
                        </div>
                    ))}
                </div>
            )}
        </Container>
    </div>
  )
}

export default MyPosts