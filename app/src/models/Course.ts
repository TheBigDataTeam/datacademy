export type Course = {
    id: number
    title: string
    description: string
    theme: string
    author: string
    authorId: number
    techstack: Array<string>
    syllabus: Array<string>
    beneficiars: Array<string>
    duration: string
    difficulty: string
    createdOn: string
    updatedOn: string
}