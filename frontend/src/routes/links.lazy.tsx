import { createLazyFileRoute } from '@tanstack/react-router'
import { Button } from "@/components/ui/button"
import { bulkAddScholarshipLinks } from '@/lib/pb';
import { useMutation } from '@tanstack/react-query';
import { useState } from 'react';

export const Route = createLazyFileRoute('/links')({
  component: LinksPage,
})

export function LinksPage() {
  const [links, setLinks] = useState('');
  const mutation = useMutation({
    mutationFn: async () => {
      await bulkAddScholarshipLinks(links);
      setLinks('');
    },
  })

  return (
    <>
      <div className="flex items-center">
        <h1 className="text-lg font-semibold md:text-2xl">Import Links</h1>
      </div>
      <textarea
        disabled={mutation.isPending}
        className="flex flex-1 items-center justify-center rounded-lg border border-dashed shadow-sm"
        onChange={(e) => setLinks(e.target.value)}
        value={links}
      />
      <Button
        className="flex items-center justify-center rounded-lg border shadow-sm"
        onClick={async () => {
          await mutation.mutateAsync();
        }}>
          Add
        </Button>
    </>
  )
}
