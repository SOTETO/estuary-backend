using System;

namespace estuary_backend.Models
{
    public class ContentLink
    {
        public Guid Id { get; set; }

        public Content Content { get; set; }

        public string LinkTag { get; set; }
    }
}
