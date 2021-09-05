
import Adapters, {TypeORMUserModel} from 'next-auth/adapters';

export class User extends TypeORMUserModel{
    constructor(name?: string,
                email?: string,
                image?: string,
                emailVerified?: Date) {
        super(name, email, image, emailVerified);

    }

}

export const UserSchema = {
    name: "User",
    target: User,
    columns: {
        ...Adapters.TypeORM.Models.User.schema.columns,
    },
}