import PocketBase, { RecordService } from 'pocketbase';

type BaseCollection<T extends Record<string, any> = {}> = {
    id: string;
    created: string;
    updated: string;
} & T

type ScholarshipUrlCollection = BaseCollection<{
    url: string;
}>

interface TypedPocketBase extends PocketBase {
    collection(idOrName: string): RecordService<BaseCollection> // default fallback for any other collection
    collection(idOrName: "scholarshipUrls"): RecordService<ScholarshipUrlCollection>
  }

const pb = new PocketBase('http://127.0.0.1:8090') as TypedPocketBase;

export const bulkAddScholarshipLinks = async (links: string) => {
    const splitLinks = links.split('\n').map(link => link.trim()).filter(link => link.length > 0);
    if (splitLinks.length === 0) {
        throw new Error('No links to import');
    }
    for (const link of splitLinks) {
        await pb.collection('scholarship_urls').create({
            url: link,
        });
    }
};
