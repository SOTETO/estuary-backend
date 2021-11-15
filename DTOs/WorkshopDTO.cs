namespace estuary_backend.DTOs
{
    public class WorkshopDTO : BaseWorkshopDTO
    {
        public string[] authors { get; set; }

        public int status { get; set; }

        public WorkshopContentDTO[] content { get; set; }
    }
}
