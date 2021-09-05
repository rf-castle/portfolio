import NextAuth from "next-auth"
import Providers from "next-auth/providers"

export default NextAuth({
    providers: [
        Providers.Discord({
            clientId: process.env.DISCORD_ID,
            clientSecret: process.env.DISCORD_SECRET,
            scope: "identify guilds"
        }),
    ],
    // SQL or MongoDB database (or leave empty)
    database: process.env.DATABASE_URL,
    callbacks: {
        async signIn(user, account, profile) {
            user.provider = account.provider
            return true
        },
        async session(session, userOrToken){
            return session
        },
        async jwt(token, user, account, profile, isNewUser) {
            return token
        }
    }
})