using System;
using System.Collections.Generic;

namespace estuary_backend.Models
{
    public class Content
    {
        public Guid Id {  get; set; }

        public string Title { get; set; }

        public Workshop Workshop { get; set; }

        public Guid WorkshopId {  get; set; }

        public List<ContentLink> RelatedContent { get; set; }

        public List<Like> Likes { get; set; }

    }
}
