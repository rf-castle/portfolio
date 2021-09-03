// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import {Page} from '../../components/pageList/pageList'

export default function handler(
    req: NextApiRequest,
    res: NextApiResponse<Page[]>
) {
    res.status(200).json([])
}
