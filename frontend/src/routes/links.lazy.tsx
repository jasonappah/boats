import { createLazyFileRoute } from '@tanstack/react-router'
import { Button } from "@/components/ui/button"
import { bulkAddScholarshipLinks } from '@/lib/pb';
import { useMutation } from '@tanstack/react-query';
import { useState } from 'react';
// TODO: update wails config to gen js to src/wails
import { FetchScholarshipHTML } from '@/../wailsjs/go/main/App';
import { ArrowDownCircle } from 'lucide-react';
import { Textarea } from '@/components/ui/textarea';

export const Route = createLazyFileRoute('/links')({
  component: LinksPage,
})

export function LinksPage() {
  const [links, setLinks] = useState('');
  const addLinksMutation = useMutation({
    mutationFn: async () => {
      await bulkAddScholarshipLinks(links);
      setLinks('');
    },
  })

  const importScholarshipsMutation = useMutation({
    mutationFn: async () => {
      await FetchScholarshipHTML();
    },
  })

  return (
    <>
      <div className="flex items-center gap-5">
        <h1 className="text-lg font-semibold md:text-2xl">Add Links</h1>
        <Button className="flex items-center justify-center rounded-lg border shadow-sm gap-2" loading={importScholarshipsMutation.isPending} onClick={async () => {
          await importScholarshipsMutation.mutateAsync();
        }}>
          <ArrowDownCircle className="h-5 w-5" />
          Import Scholarships
        </Button>
      </div>
      <Textarea
        disabled={addLinksMutation.isPending}
        className="flex flex-1 items-center justify-center rounded-lg border border-dashed shadow-sm"
        onChange={(e) => setLinks(e.target.value)}
        value={links}
      />
      <Button
        loading={addLinksMutation.isPending}
        className="flex items-center justify-center rounded-lg border shadow-sm"
        onClick={async () => {
          await addLinksMutation.mutateAsync();
        }}>
        Add
      </Button>
    </>
  )
}
